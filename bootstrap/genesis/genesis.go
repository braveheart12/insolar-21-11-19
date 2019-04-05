//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package genesis

import (
	"context"
	"crypto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/pkg/errors"

	"github.com/insolar/insolar/application/contract/member"
	"github.com/insolar/insolar/application/contract/nodedomain"
	"github.com/insolar/insolar/application/contract/noderecord"
	"github.com/insolar/insolar/application/contract/rootdomain"
	"github.com/insolar/insolar/application/contract/wallet"
	"github.com/insolar/insolar/certificate"
	"github.com/insolar/insolar/insolar"
	"github.com/insolar/insolar/insolar/message"
	"github.com/insolar/insolar/instrumentation/inslogger"
	"github.com/insolar/insolar/internal/ledger/artifact"
	"github.com/insolar/insolar/ledger/storage/genesis"
	"github.com/insolar/insolar/platformpolicy"
)

const (
	nodeDomain        = "nodedomain"
	nodeRecord        = "noderecord"
	rootDomain        = "rootdomain"
	walletContract    = "wallet"
	memberContract    = "member"
	allowanceContract = "allowance"
	nodeAmount        = 32
)

var contractNames = []string{walletContract, memberContract, allowanceContract, rootDomain, nodeDomain, nodeRecord}

type nodeInfo struct {
	privateKey crypto.PrivateKey
	publicKey  string
	ref        *insolar.Reference
}

// Generator is a component for generating RootDomain instance and genesis contracts.
type Generator struct {
	ArtifactManager artifact.Manager `inject:""`
	GenesisState    genesis.State    `inject:""`

	config *Config

	rootDomainRef *insolar.Reference
	nodeDomainRef *insolar.Reference
	rootMemberRef *insolar.Reference

	keyOut string
}

// NewGenerator creates new Generator.
func NewGenerator(genesisConfigPath string, genesisKeyOut string) (*Generator, error) {
	config, err := ParseGenesisConfig(genesisConfigPath)
	genesis := &Generator{
		rootDomainRef: &insolar.Reference{},
		config:        config,
		keyOut:        genesisKeyOut,
	}
	return genesis, err
}

func buildSmartContracts(ctx context.Context, cb *ContractsBuilder, rootDomainID *insolar.ID) error {
	inslog := inslogger.FromContext(ctx)
	inslog.Info("[ buildSmartContracts ] building contracts:", contractNames)
	contracts, err := getContractsMap()
	if err != nil {
		return errors.Wrap(err, "[ buildSmartContracts ] failed to get contracts map")
	}

	inslog.Info("[ buildSmartContracts ] Start building contracts ...")
	err = cb.Build(ctx, contracts, rootDomainID)
	if err != nil {
		return errors.Wrap(err, "[ buildSmartContracts ] couldn't build contracts")
	}
	inslog.Info("[ buildSmartContracts ] Stop building contracts ...")

	return nil
}

func (g *Generator) activateRootDomain(
	ctx context.Context,
	cb *ContractsBuilder,
	contractID *insolar.ID,
) (artifact.ObjectDescriptor, error) {
	rd, err := rootdomain.NewRootDomain()
	if err != nil {
		return nil, errors.Wrap(err, "[ ActivateRootDomain ]")
	}

	instanceData, err := insolar.Serialize(rd)
	if err != nil {
		return nil, errors.Wrap(err, "[ ActivateRootDomain ]")
	}

	contract := insolar.NewReference(*contractID, *contractID)
	desc, err := g.ArtifactManager.ActivateObject(
		ctx,
		insolar.Reference{},
		*contract,
		*g.GenesisState.GenesisRef(),
		*cb.Prototypes[rootDomain],
		false,
		instanceData,
	)
	if err != nil {
		return nil, errors.Wrap(err, "[ ActivateRootDomain ] Couldn't create rootdomain instance")
	}
	_, err = g.ArtifactManager.RegisterResult(ctx, *g.GenesisState.GenesisRef(), *contract, nil)
	if err != nil {
		return nil, errors.Wrap(err, "[ ActivateRootDomain ] Couldn't create rootdomain instance")
	}
	g.rootDomainRef = contract

	return desc, nil
}

func (g *Generator) activateNodeDomain(
	ctx context.Context, domain *insolar.ID, cb *ContractsBuilder,
) (artifact.ObjectDescriptor, error) {
	nd, _ := nodedomain.NewNodeDomain()

	instanceData, err := insolar.Serialize(nd)
	if err != nil {
		return nil, errors.Wrap(err, "[ ActivateNodeDomain ] node domain serialization")
	}

	contractID, err := g.ArtifactManager.RegisterRequest(
		ctx,
		*g.rootDomainRef,
		&message.Parcel{
			Msg: &message.GenesisRequest{Name: "NodeDomain"},
		},
	)

	if err != nil {
		return nil, errors.Wrap(err, "[ ActivateNodeDomain ] couldn't create nodedomain instance")
	}
	contract := insolar.NewReference(*domain, *contractID)
	desc, err := g.ArtifactManager.ActivateObject(
		ctx,
		insolar.Reference{},
		*contract,
		*g.rootDomainRef,
		*cb.Prototypes[nodeDomain],
		false,
		instanceData,
	)
	if err != nil {
		return nil, errors.Wrap(err, "[ ActivateNodeDomain ] couldn't create nodedomain instance")
	}
	_, err = g.ArtifactManager.RegisterResult(ctx, *g.rootDomainRef, *contract, nil)
	if err != nil {
		return nil, errors.Wrap(err, "[ ActivateNodeDomain ] couldn't create nodedomain instance")
	}

	g.nodeDomainRef = contract

	return desc, nil
}

func (g *Generator) activateRootMember(
	ctx context.Context, domain *insolar.ID, cb *ContractsBuilder, rootPubKey string,
) error {

	m, err := member.New("RootMember", rootPubKey)
	if err != nil {
		return errors.Wrap(err, "[ ActivateRootMember ]")
	}

	instanceData, err := insolar.Serialize(m)
	if err != nil {
		return errors.Wrap(err, "[ ActivateRootMember ]")
	}

	contractID, err := g.ArtifactManager.RegisterRequest(
		ctx,
		*g.rootDomainRef,
		&message.Parcel{
			Msg: &message.GenesisRequest{Name: "RootMember"},
		},
	)

	if err != nil {
		return errors.Wrap(err, "[ ActivateRootMember ] couldn't create root member instance")
	}
	contract := insolar.NewReference(*domain, *contractID)
	_, err = g.ArtifactManager.ActivateObject(
		ctx,
		insolar.Reference{},
		*contract,
		*g.rootDomainRef,
		*cb.Prototypes[memberContract],
		false,
		instanceData,
	)
	if err != nil {
		return errors.Wrap(err, "[ ActivateRootMember ] couldn't create root member instance")
	}
	_, err = g.ArtifactManager.RegisterResult(ctx, *g.rootDomainRef, *contract, nil)
	if err != nil {
		return errors.Wrap(err, "[ ActivateRootMember ] couldn't create root member instance")
	}
	g.rootMemberRef = contract
	return nil
}

// TODO: this is not required since we refer by request id.
func (g *Generator) updateRootDomain(
	ctx context.Context, domainDesc artifact.ObjectDescriptor,
) error {
	updateData, err := insolar.Serialize(&rootdomain.RootDomain{RootMember: *g.rootMemberRef, NodeDomainRef: *g.nodeDomainRef})
	if err != nil {
		return errors.Wrap(err, "[ updateRootDomain ]")
	}
	_, err = g.ArtifactManager.UpdateObject(
		ctx,
		insolar.Reference{},
		insolar.Reference{},
		domainDesc,
		updateData,
	)
	if err != nil {
		return errors.Wrap(err, "[ updateRootDomain ]")
	}

	return nil
}

func (g *Generator) activateRootMemberWallet(
	ctx context.Context, domain *insolar.ID, cb *ContractsBuilder,
) error {

	w, err := wallet.New(g.config.RootBalance)
	if err != nil {
		return errors.Wrap(err, "[ ActivateRootWallet ]")
	}

	instanceData, err := insolar.Serialize(w)
	if err != nil {
		return errors.Wrap(err, "[ ActivateRootWallet ]")
	}

	contractID, err := g.ArtifactManager.RegisterRequest(
		ctx,
		*g.rootDomainRef,
		&message.Parcel{
			Msg: &message.GenesisRequest{Name: "RootWallet"},
		},
	)

	if err != nil {
		return errors.Wrap(err, "[ ActivateRootWallet ] couldn't create root wallet")
	}
	contract := insolar.NewReference(*domain, *contractID)
	_, err = g.ArtifactManager.ActivateObject(
		ctx,
		insolar.Reference{},
		*contract,
		*g.rootMemberRef,
		*cb.Prototypes[walletContract],
		true,
		instanceData,
	)
	if err != nil {
		return errors.Wrap(err, "[ ActivateRootWallet ] couldn't create root wallet")
	}
	_, err = g.ArtifactManager.RegisterResult(ctx, *g.rootDomainRef, *contract, nil)
	if err != nil {
		return errors.Wrap(err, "[ ActivateRootWallet ] couldn't create root wallet")
	}

	return nil
}

func (g *Generator) activateSmartContracts(
	ctx context.Context, cb *ContractsBuilder, rootPubKey string, rootDomainID *insolar.ID,
) ([]genesisNode, error) {

	rootDomainDesc, err := g.activateRootDomain(ctx, cb, rootDomainID)
	errMsg := "[ ActivateSmartContracts ]"
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	nodeDomainDesc, err := g.activateNodeDomain(ctx, rootDomainID, cb)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	err = g.activateRootMember(ctx, rootDomainID, cb, rootPubKey)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	// TODO: this is not required since we refer by request id.
	err = g.updateRootDomain(ctx, rootDomainDesc)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	err = g.activateRootMemberWallet(ctx, rootDomainID, cb)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	indexMap := make(map[string]string)

	discoveryNodes, indexMap, err := g.addDiscoveryIndex(ctx, cb, indexMap)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)

	}
	indexMap, err = g.addIndex(ctx, cb, indexMap)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)

	}

	err = g.updateNodeDomainIndex(ctx, nodeDomainDesc, indexMap)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}

	return discoveryNodes, nil
}

type genesisNode struct {
	node    certificate.BootstrapNode
	privKey crypto.PrivateKey
	ref     *insolar.Reference
	role    string
}

func (g *Generator) activateDiscoveryNodes(ctx context.Context, cb *ContractsBuilder, nodesInfo []nodeInfo) ([]genesisNode, error) {
	if len(nodesInfo) != len(g.config.DiscoveryNodes) {
		return nil, errors.New("[ activateDiscoveryNodes ] len of nodesInfo param must be equal to len of DiscoveryNodes in genesis config")
	}

	nodes := make([]genesisNode, len(g.config.DiscoveryNodes))

	for i, discoverNode := range g.config.DiscoveryNodes {
		privKey := nodesInfo[i].privateKey
		nodePubKey := nodesInfo[i].publicKey

		nodeState := &noderecord.NodeRecord{
			Record: noderecord.RecordInfo{
				PublicKey: nodePubKey,
				Role:      insolar.GetStaticRoleFromString(discoverNode.Role),
			},
		}
		contract, err := g.activateNodeRecord(ctx, cb, nodeState, "discoverynoderecord_"+strconv.Itoa(i))
		if err != nil {
			return nil, errors.Wrap(err, "[ activateDiscoveryNodes ] Couldn't activateNodeRecord node instance")
		}

		nodes[i] = genesisNode{
			node: certificate.BootstrapNode{
				PublicKey: nodePubKey,
				Host:      discoverNode.Host,
				NodeRef:   contract.String(),
			},
			privKey: privKey,
			ref:     contract,
			role:    discoverNode.Role,
		}
	}
	return nodes, nil
}

func (g *Generator) activateNodes(ctx context.Context, cb *ContractsBuilder, nodes []nodeInfo) ([]nodeInfo, error) {
	var updatedNodes []nodeInfo

	for i, node := range nodes {
		nodeState := &noderecord.NodeRecord{
			Record: noderecord.RecordInfo{
				PublicKey: node.publicKey,
				Role:      insolar.StaticRoleVirtual,
			},
		}
		contract, err := g.activateNodeRecord(ctx, cb, nodeState, "noderecord_"+strconv.Itoa(i))
		if err != nil {
			return nil, errors.Wrap(err, "[ activateNodes ] Couldn't activateNodeRecord node instance")
		}
		updatedNode := nodeInfo{
			ref:       contract,
			publicKey: node.publicKey,
		}
		updatedNodes = append(updatedNodes, updatedNode)
	}

	return updatedNodes, nil
}

func (g *Generator) activateNodeRecord(ctx context.Context, cb *ContractsBuilder, record *noderecord.NodeRecord, name string) (*insolar.Reference, error) {
	nodeData, err := insolar.Serialize(record)
	if err != nil {
		return nil, errors.Wrap(err, "[ activateNodeRecord ] Couldn't serialize node instance")
	}

	nodeID, err := g.ArtifactManager.RegisterRequest(
		ctx,
		*g.rootDomainRef,
		&message.Parcel{
			Msg: &message.GenesisRequest{Name: name},
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "[ activateNodeRecord ] Couldn't register request")
	}
	contract := insolar.NewReference(*g.rootDomainRef.Record(), *nodeID)
	_, err = g.ArtifactManager.ActivateObject(
		ctx,
		insolar.Reference{},
		*contract,
		*g.nodeDomainRef,
		*cb.Prototypes[nodeRecord],
		false,
		nodeData,
	)
	if err != nil {
		return nil, errors.Wrap(err, "[ activateNodeRecord ] Could'n activateNodeRecord node object")
	}
	_, err = g.ArtifactManager.RegisterResult(ctx, *g.rootDomainRef, *contract, nil)
	if err != nil {
		return nil, errors.Wrap(err, "[ activateNodeRecord ] Couldn't register result")
	}
	return contract, nil
}

func (g *Generator) addDiscoveryIndex(ctx context.Context, cb *ContractsBuilder, indexMap map[string]string) ([]genesisNode, map[string]string, error) {
	errMsg := "[ addDiscoveryIndex ]"
	discoveryKeysPath, err := absPath(g.config.DiscoveryKeysDir)
	if err != nil {
		return nil, nil, errors.Wrap(err, errMsg)
	}
	discoveryKeys, err := g.uploadKeys(ctx, discoveryKeysPath, len(g.config.DiscoveryNodes))
	if err != nil {
		return nil, nil, errors.Wrap(err, errMsg)
	}
	discoveryNodes, err := g.activateDiscoveryNodes(ctx, cb, discoveryKeys)
	if err != nil {
		return nil, nil, errors.Wrap(err, errMsg)
	}
	for _, node := range discoveryNodes {
		indexMap[node.node.PublicKey] = node.ref.String()
	}
	return discoveryNodes, indexMap, nil
}

func (g *Generator) addIndex(ctx context.Context, cb *ContractsBuilder, indexMap map[string]string) (map[string]string, error) {
	errMsg := "[ addIndex ]"
	nodeKeysPath, err := absPath(g.config.NodeKeysDir)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	userKeys, err := g.uploadKeys(ctx, nodeKeysPath, nodeAmount)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	nodes, err := g.activateNodes(ctx, cb, userKeys)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	for _, node := range nodes {
		indexMap[node.publicKey] = node.ref.String()
	}
	return indexMap, nil
}

func (g *Generator) createKeys(ctx context.Context, path string, amount int) error {
	err := os.RemoveAll(path)
	if err != nil {
		return errors.Wrap(err, "[ createKeys ] couldn't remove old dir")
	}

	for i := 0; i < amount; i++ {
		ks := platformpolicy.NewKeyProcessor()

		privKey, err := ks.GeneratePrivateKey()
		if err != nil {
			return errors.Wrap(err, "[ createKeys ] couldn't generate private key")
		}

		privKeyStr, err := ks.ExportPrivateKeyPEM(privKey)
		if err != nil {
			return errors.Wrap(err, "[ createKeys ] couldn't export private key")
		}

		pubKeyStr, err := ks.ExportPublicKeyPEM(ks.ExtractPublicKey(privKey))
		if err != nil {
			return errors.Wrap(err, "[ createKeys ] couldn't export public key")
		}

		result, err := json.MarshalIndent(map[string]interface{}{
			"private_key": string(privKeyStr),
			"public_key":  string(pubKeyStr),
		}, "", "    ")
		if err != nil {
			return errors.Wrap(err, "[ createKeys ] couldn't marshal keys")
		}

		name := fmt.Sprintf(g.config.KeysNameFormat, i)
		err = WriteFile(path, name, string(result))
		if err != nil {
			return errors.Wrap(err, "[ createKeys ] couldn't write keys to file")
		}
	}

	return nil
}

func (g *Generator) uploadKeys(ctx context.Context, path string, amount int) ([]nodeInfo, error) {
	var err error
	if !g.config.ReuseKeys {
		err = g.createKeys(ctx, path, amount)
		if err != nil {
			return nil, errors.Wrap(err, "[ uploadKeys ]")
		}
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.Wrap(err, "[ uploadKeys ] can't read dir")
	}
	if len(files) != amount {
		return nil, errors.New(fmt.Sprintf("[ uploadKeys ] amount of nodes != amount of files in directory: %d != %d", len(files), amount))
	}

	var keys []nodeInfo
	for _, f := range files {
		privKey, nodePubKey, err := getKeysFromFile(ctx, filepath.Join(path, f.Name()))
		if err != nil {
			return nil, errors.Wrap(err, "[ uploadKeys ] can't get keys from file")
		}

		key := nodeInfo{
			publicKey:  nodePubKey,
			privateKey: privKey,
		}
		keys = append(keys, key)
	}

	return keys, nil
}

// Start creates types and RootDomain instance
func (g *Generator) Start(ctx context.Context) error {
	inslog := inslogger.FromContext(ctx)
	inslog.Info("[ Genesis ] Starting  ...")
	defer inslog.Info("[ Genesis ] Finished.")

	rootDomainID, err := g.ArtifactManager.RegisterRequest(
		ctx,
		*g.GenesisState.GenesisRef(),
		&message.Parcel{
			Msg: &message.GenesisRequest{
				Name: rootDomain,
			},
		},
	)
	if err != nil {
		panic(errors.Wrap(err, "[ Genesis ] Couldn't create rootdomain instance"))
	}

	inslog.Info("[ Genesis ] NewContractBuilder ...")
	cb := NewContractBuilder(*g.GenesisState.GenesisRef(), g.ArtifactManager)
	defer cb.Clean()

	inslog.Info("[ Genesis ] buildSmartContracts ...")
	err = buildSmartContracts(ctx, cb, rootDomainID)
	if err != nil {
		panic(errors.Wrap(err, "[ Genesis ] couldn't build contracts"))
	}

	inslog.Info("[ Genesis ] getKeysFromFile ...")
	_, rootPubKey, err := getKeysFromFile(ctx, g.config.RootKeysFile)
	if err != nil {
		return errors.Wrap(err, "[ Genesis ] couldn't get root keys")
	}

	inslog.Info("[ Genesis ] activateSmartContracts ...")
	nodes, err := g.activateSmartContracts(ctx, cb, rootPubKey, rootDomainID)
	if err != nil {
		panic(errors.Wrap(err, "[ Genesis ] could't activate smart contracts"))
	}

	inslog.Info("[ Genesis ] makeCertificates ...")
	err = g.makeCertificates(nodes)
	if err != nil {
		return errors.Wrap(err, "[ Genesis ] Couldn't generate discovery certificates")
	}

	return nil
}

func (g *Generator) makeCertificates(nodes []genesisNode) error {
	certs := make([]certificate.Certificate, len(nodes))
	for i, node := range nodes {
		certs[i].Role = node.role
		certs[i].Reference = node.ref.String()
		certs[i].PublicKey = node.node.PublicKey
		certs[i].RootDomainReference = g.rootDomainRef.String()
		certs[i].MajorityRule = g.config.MajorityRule
		certs[i].MinRoles.Virtual = g.config.MinRoles.Virtual
		certs[i].MinRoles.HeavyMaterial = g.config.MinRoles.HeavyMaterial
		certs[i].MinRoles.LightMaterial = g.config.MinRoles.LightMaterial
		certs[i].BootstrapNodes = make([]certificate.BootstrapNode, len(nodes))
		for j, node := range nodes {
			certs[i].BootstrapNodes[j] = node.node
		}
	}

	var err error
	for i := range nodes {
		for j, node := range nodes {
			certs[i].BootstrapNodes[j].NetworkSign, err = certs[i].SignNetworkPart(node.privKey)
			if err != nil {
				return errors.Wrapf(err, "[ makeCertificates ] Can't SignNetworkPart for %s", node.ref.String())
			}

			certs[i].BootstrapNodes[j].NodeSign, err = certs[i].SignNodePart(node.privKey)
			if err != nil {
				return errors.Wrapf(err, "[ makeCertificates ] Can't SignNodePart for %s", node.ref.String())
			}
		}

		// save cert to disk
		cert, err := json.MarshalIndent(certs[i], "", "  ")
		if err != nil {
			return errors.Wrapf(err, "[ makeCertificates ] Can't MarshalIndent")
		}

		if len(g.config.DiscoveryNodes[i].CertName) == 0 {
			return errors.New("[ makeCertificates ] cert_name must not be empty for node " + strconv.Itoa(i+1))
		}

		err = ioutil.WriteFile(path.Join(g.keyOut, g.config.DiscoveryNodes[i].CertName), cert, 0644)
		if err != nil {
			return errors.Wrap(err, "[ makeCertificates ] WriteFile")
		}
	}
	return nil
}

func (g *Generator) updateNodeDomainIndex(ctx context.Context, nodeDomainDesc artifact.ObjectDescriptor, indexMap map[string]string) error {
	updateData, err := insolar.Serialize(
		&nodedomain.NodeDomain{
			NodeIndexPK: indexMap,
		},
	)
	if err != nil {
		return errors.Wrap(err, "[ updateNodeDomainIndex ]  Couldn't serialize NodeDomain")
	}

	_, err = g.ArtifactManager.UpdateObject(
		ctx,
		*g.rootDomainRef,
		*g.nodeDomainRef,
		nodeDomainDesc,
		updateData,
	)
	if err != nil {
		return errors.Wrap(err, "[ updateNodeDomainIndex ]  Couldn't update NodeDomain")
	}

	return nil
}
