package commit

import (
	"github.com/ssvlabs/ssv-spec/qbft/spectest/tests"
	"github.com/ssvlabs/ssv-spec/types"
	"github.com/ssvlabs/ssv-spec/types/testingutils"
)

// PostCutoff tests processing a commit msg when round >= cutoff
func PostCutoff() tests.SpecTest {
	ks := testingutils.Testing4SharesSet()

	pre := testingutils.BaseInstance()
	pre.State.Round = 15

	msgs := []*types.SignedSSVMessage{
		testingutils.TestingCommitMessageWithRound(ks.OperatorKeys[1], types.OperatorID(1), 15),
	}

	return &tests.MsgProcessingSpecTest{
		Name:          "round cutoff commit message",
		Pre:           pre,
		InputMessages: msgs,
		ExpectedError: "instance stopped processing messages",
	}
}
