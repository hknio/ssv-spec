package spectest

import (
	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests"
	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests/commit"
	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests/controller/decided"
	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests/controller/futuremsg"
	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests/controller/latemsg"
	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests/controller/processmsg"
	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests/messages"
	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests/prepare"
	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests/proposal"
	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests/proposer"
	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests/roundchange"
	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests/startinstance"
	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests/timeout"
)

var AllTests = []tests.TestF{
	timeout.Round1,
	timeout.Round2,
	timeout.Round3,
	timeout.Round5,
	timeout.Round15,

	decided.Valid,
	decided.HasQuorum,
	decided.LateDecided,
	decided.LateDecidedBiggerQuorum,
	decided.LateDecidedSmallerQuorum,
	decided.NoQuorum,
	decided.DuplicateMsg,
	decided.DuplicateSigners,
	decided.Invalid,
	decided.InvalidFullData,
	decided.InvalidValCheckData,
	decided.PastInstance,
	decided.UnknownSigner,
	decided.WrongMsgType,
	decided.WrongSignature,
	decided.MultiDecidedInstances,
	decided.FutureInstance,
	decided.CurrentInstance,
	decided.CurrentInstancePastRound,
	decided.CurrentInstanceFutureRound,

	processmsg.MsgError,
	processmsg.BroadcastedDecided,
	processmsg.SingleConsensusMsg,
	processmsg.FullDecided,
	processmsg.InvalidIdentifier,
	processmsg.NoInstanceRunning,

	latemsg.LateCommit,
	latemsg.LateCommitPastRound,
	latemsg.LateCommitPastInstance,
	latemsg.LatePrepare,
	latemsg.LatePreparePastInstance,
	latemsg.LatePreparePastRound,
	latemsg.LateProposal,
	latemsg.LateProposalPastInstance,
	latemsg.LateProposalPastRound,
	latemsg.LateProposalPastInstanceNonDuplicate,
	latemsg.LateRoundChange,
	latemsg.LateRoundChangePastInstance,
	latemsg.LateRoundChangePastRound,
	latemsg.FullFlowAfterDecided,

	futuremsg.ValidMsg,

	startinstance.Valid,
	startinstance.EmptyValue,
	startinstance.NilValue,
	startinstance.PostFutureDecided,
	startinstance.FirstHeight,
	startinstance.PreviousDecided,
	startinstance.PreviousNotDecided,
	startinstance.InvalidValue,
	startinstance.EqualHeightRunningInstance,
	startinstance.EqualHeightNoRunningInstance,
	startinstance.LowerHeight,

	proposer.FourOperators,
	proposer.SevenOperators,
	proposer.TenOperators,
	proposer.ThirteenOperators,

	messages.RoundChangePrePreparedJustifications,
	messages.RoundChangeNotPreparedJustifications,
	messages.CommitDataEncoding,
	messages.MsgNilIdentifier,
	messages.MsgNonZeroIdentifier,
	messages.MsgTypeUnknown,
	messages.PrepareDataEncoding,
	messages.PrepareJustificationsUnmarshalling,
	messages.ProposeDataEncoding,
	messages.SignedMsgNoSigners,
	messages.SignedMsgDuplicateSigners,

	messages.SignedMsgMultiSigners,
	messages.GetRoot,
	messages.RoundChangeJustificationsUnmarshalling,
	messages.RoundChangePrepared,
	messages.SignedMessageEncoding,
	messages.CreateProposal,
	messages.CreateProposalPreviouslyPrepared,
	messages.CreateProposalNotPreviouslyPrepared,
	messages.CreatePrepare,
	messages.CreateCommit,
	messages.CreateRoundChange,
	messages.InvalidHashDataRoot,
	messages.InvalidPrepareJustificationsUnmarshalling,
	messages.InvalidRoundChangeJustificationsUnmarshalling,
	messages.CreateRoundChangePreviouslyPrepared,
	messages.CreateRoundChangeNoJustificationQuorum,
	messages.RoundChangeDataEncoding,
	messages.SignedMessageSigner0,
	messages.SSZMarshaling,
	messages.MarshalJustificationsWithoutFullData,
	messages.MarshalJustificationsWithFullData,
	messages.MarshalJustifications,
	messages.UnmarshalJustifications,
	messages.ValidHashDataRoot,

	tests.HappyFlow,
	tests.SevenOperators,
	tests.TenOperators,
	tests.ThirteenOperators,

	proposal.InvalidFullData,
	proposal.PastRoundProposalPrevPrepared,
	proposal.NotPreparedPreviouslyJustification,
	proposal.PreparedPreviouslyJustification,
	proposal.DifferentJustifications,
	proposal.JustificationsNotHeighest,
	proposal.DuplicateMsg,
	proposal.DuplicateMsgDifferentRoot,
	proposal.FirstRoundJustification,
	proposal.FutureRoundPrevNotPrepared,
	proposal.FutureRound,
	proposal.InvalidRoundChangeJustificationPrepared,
	proposal.InvalidRoundChangeJustification,
	proposal.PreparedPreviouslyNoRCJustificationQuorum,
	proposal.NoRCJustification,
	proposal.PreparedPreviouslyNoPrepareJustificationQuorum,
	proposal.PreparedPreviouslyDuplicatePrepareMsg,
	proposal.PreparedPreviouslyDuplicatePrepareQuorum,
	proposal.PreparedPreviouslyDuplicateRCMsg,
	proposal.PreparedPreviouslyDuplicateRCQuorum,
	proposal.DuplicateRCMsg,
	proposal.InvalidPrepareJustificationValue,
	proposal.InvalidPrepareJustificationRound,
	proposal.InvalidValueCheck,
	proposal.MultiSigner,
	proposal.PostDecided,
	proposal.PostPrepared,
	proposal.SecondProposalForRound,
	proposal.WrongHeight,
	proposal.WrongProposer,
	proposal.UnknownSigner,
	proposal.ForceStop,
	proposal.PostCutoff,

	prepare.DuplicateMsg,
	prepare.HappyFlow,
	prepare.MultiSigner,
	prepare.NoPreviousProposal,
	prepare.OldRound,
	prepare.FutureRound,
	prepare.PostDecided,
	prepare.WrongData,
	prepare.WrongHeight,
	prepare.UnknownSigner,
	prepare.PrepareQuorumTriggeredTwice,
	prepare.PrepareQuorumTriggeredTwiceLateCommit,
	prepare.ForceStop,
	prepare.PostCutoff,

	commit.CurrentRound,
	commit.FutureRound,
	commit.PastRound,
	commit.DuplicateMsg,
	commit.HappyFlow,
	commit.PostDecided,
	commit.WrongData1,
	commit.WrongData2,
	commit.MultiSignerWithOverlap,
	commit.MultiSignerNoOverlap,
	commit.DuplicateSigners,
	commit.NoPrevAcceptedProposal,
	commit.WrongHeight,
	commit.UnknownSigner,
	commit.InvalidValCheck,
	commit.NoPrepareQuorum,
	commit.NoCommitQuorum,
	commit.ForceStop,
	commit.PostCutoff,

	roundchange.HappyFlow,
	roundchange.WrongHeight,
	roundchange.UnknownSigner,
	roundchange.MultiSigner,
	roundchange.QuorumNotPrepared,
	roundchange.QuorumPrepared,
	roundchange.PeerPrepared,
	roundchange.PeerPreparedDifferentHeights,
	roundchange.JustificationWrongRound,
	roundchange.JustificationNoQuorum,
	roundchange.JustificationMultiSigners,
	roundchange.JustificationInvalidSig,
	roundchange.JustificationInvalidRound,
	roundchange.JustificationDuplicateMsg,
	roundchange.QuorumNotTriggeredTwice,
	roundchange.QuorumNotTriggeredTwiceJustificationIgnored,
	roundchange.F1DifferentFutureRounds,
	roundchange.F1DifferentFutureRoundsNotPrepared,
	roundchange.PastRound,
	roundchange.DuplicateMsgQuorum,
	roundchange.DuplicateMsgQuorumPreparedRCFirst,
	roundchange.DuplicateMsg,
	roundchange.NotProposer,
	roundchange.ValidJustification,
	roundchange.F1DuplicateSigner,
	roundchange.F1DuplicateSignerNotPrepared,
	roundchange.F1Speedup,
	roundchange.F1SpeedupPrevPrepared,
	roundchange.AfterProposal,
	roundchange.RoundChangePartialQuorum,
	roundchange.QuorumOrder2,
	roundchange.QuorumOrder1,
	roundchange.QuorumMsgNotPrepared,
	roundchange.JustificationPastRound,
	roundchange.ForceStop,
	roundchange.PostCutoff,
	roundchange.EmptySigners,
	roundchange.NonUniqueSigner,
	roundchange.ZeroSigner,
}
