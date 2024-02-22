package method

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
)

var MethodsMap = make(map[abi.MethodNum]string)

func init() {

	methodsMiner := builtin.MethodsMiner

	MethodsMap[methodsMiner.Constructor] = "Constructor"
	MethodsMap[methodsMiner.ControlAddresses] = "ControlAddresses"
	MethodsMap[methodsMiner.ChangeWorkerAddress] = "ChangeWorkerAddress"
	MethodsMap[methodsMiner.ChangeWorkerAddressExported] = "ChangeWorkerAddressExported"
	MethodsMap[methodsMiner.ChangePeerID] = "ChangePeerID"
	MethodsMap[methodsMiner.ChangePeerIDExported] = "ChangePeerIDExported"
	MethodsMap[methodsMiner.SubmitWindowedPoSt] = "SubmitWindowedPoSt"
	MethodsMap[methodsMiner.PreCommitSector] = "PreCommitSector"
	MethodsMap[methodsMiner.ProveCommitSector] = "ProveCommitSector"
	MethodsMap[methodsMiner.ExtendSectorExpiration] = "ExtendSectorExpiration"
	MethodsMap[methodsMiner.TerminateSectors] = "TerminateSectors"
	MethodsMap[methodsMiner.DeclareFaults] = "DeclareFaults"
	MethodsMap[methodsMiner.DeclareFaultsRecovered] = "DeclareFaultsRecovered"
	MethodsMap[methodsMiner.OnDeferredCronEvent] = "OnDeferredCronEvent"
	MethodsMap[methodsMiner.CheckSectorProven] = "CheckSectorProven"
	MethodsMap[methodsMiner.ApplyRewards] = "ApplyRewards"
	MethodsMap[methodsMiner.ReportConsensusFault] = "ReportConsensusFault"
	MethodsMap[methodsMiner.WithdrawBalance] = "WithdrawBalance"
	MethodsMap[methodsMiner.WithdrawBalanceExported] = "WithdrawBalanceExported"
	MethodsMap[methodsMiner.ConfirmSectorProofsValid] = "ConfirmSectorProofsValid"
	MethodsMap[methodsMiner.ChangeMultiaddrs] = "ChangeMultiaddrs"
	MethodsMap[methodsMiner.ChangeMultiaddrsExported] = "ChangeMultiaddrsExported"
	MethodsMap[methodsMiner.CompactPartitions] = "CompactPartitions"
	MethodsMap[methodsMiner.CompactSectorNumbers] = "CompactSectorNumbers"
	MethodsMap[methodsMiner.ConfirmChangeWorkerAddress] = "ConfirmChangeWorkerAddress"
	MethodsMap[methodsMiner.ConfirmChangeWorkerAddressExported] = "ConfirmChangeWorkerAddressExported"
	MethodsMap[methodsMiner.RepayDebt] = "RepayDebt"
	MethodsMap[methodsMiner.RepayDebtExported] = "RepayDebtExported"
	MethodsMap[methodsMiner.ChangeOwnerAddress] = "ChangeOwnerAddress"
	MethodsMap[methodsMiner.ChangeOwnerAddressExported] = "ChangeOwnerAddressExported"
	MethodsMap[methodsMiner.DisputeWindowedPoSt] = "DisputeWindowedPoSt"
	MethodsMap[methodsMiner.PreCommitSectorBatch] = "PreCommitSectorBatch"
	MethodsMap[methodsMiner.ProveCommitAggregate] = "ProveCommitAggregate"
	MethodsMap[methodsMiner.ProveReplicaUpdates] = "ProveReplicaUpdates"
	MethodsMap[methodsMiner.PreCommitSectorBatch2] = "PreCommitSectorBatch2"
	MethodsMap[methodsMiner.ProveReplicaUpdates2] = "ProveReplicaUpdates2"
	MethodsMap[methodsMiner.ChangeBeneficiary] = "ChangeBeneficiary"
	MethodsMap[methodsMiner.ChangeBeneficiaryExported] = "ChangeBeneficiaryExported"
	MethodsMap[methodsMiner.GetBeneficiary] = "GetBeneficiary"
	MethodsMap[methodsMiner.ExtendSectorExpiration2] = "ExtendSectorExpiration2"
	MethodsMap[methodsMiner.GetOwnerExported] = "GetOwnerExported"
	MethodsMap[methodsMiner.IsControllingAddressExported] = "IsControllingAddressExported"
	MethodsMap[methodsMiner.GetSectorSizeExported] = "GetSectorSizeExported"
	MethodsMap[methodsMiner.GetAvailableBalanceExported] = "GetAvailableBalanceExported"
	MethodsMap[methodsMiner.GetVestingFundsExported] = "GetVestingFundsExported"
	MethodsMap[methodsMiner.GetPeerIDExported] = "GetPeerIDExported"
	MethodsMap[methodsMiner.GetMultiaddrsExported] = "GetMultiaddrsExported"

}
