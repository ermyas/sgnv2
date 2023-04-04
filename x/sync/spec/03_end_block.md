<!--
order: 3
-->

# End-Block

End block calls are executed to tally votes and apply pending update proposals. For each pending update, if more than 2/3 voting power have voted yes, the update will be applied by the relevant module and then deleted. Otherwise, if the voting window of this pending update has close, the pending update will also be deleted.

[Implementation Reference](https://github.com/celer-network/sgnv2/blob/f9f76fb10d/x/sync/abci.go#L12-L48)