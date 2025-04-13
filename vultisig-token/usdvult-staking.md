---
description: This describes $VULT staking mechanism
---

# $VULT Staking

Users can stake $VULT tokens in the Vultisig Staking contract on Ethereum mainnet to earn rewards from ecosystem revenue.

### Key Features

* Inspiration by Sushiswap MasterChef staking
* **Revenue Sources**: Rewards come from ecosystem integrations including bridge, swap and marketplace fees
* **Fair Distribution**: Decay-based distribution mechanism prevents frontrunning and rewards long-term stakers
* **Flexible Participation**: No lockup periods—stake or unstake anytime

### Reward Distribution Mechanism

The decay-based distribution works as follows:

1. Rewards enter the staking contract and distribute gradually based on configured parameters
2. Distribution uses two key variables:
   * **Decay Interval**: Time between distributions (e.g., 1 day)
   * **Decay Factor**: Percentage distributed each interval (e.g., 10%)

{% hint style="info" %}
The vesting configuration may be updated periodically to ensure that, as revenue grows, the distribution mechanics adapt to effectively channel the increased value to stakers
{% endhint %}

### Example Distribution

With a 1-day decay interval and 10% decay factor, a 1000 USDC reward would distribute:

* Day 1: 100 USDC (10% of 1000)
* Day 2: 90 USDC (10% of remaining 900)
* Day 3: 81 USDC (10% of remaining 810)
* And so on until fully distributed

This system accommodates additional rewards during ongoing distributions, smoothing out fluctuations between high and low fee periods.

### Launch Configuration:

**TBD**

### Governance

$VULT stakers will possibly be able to vote on the Decay Interval and Decay Factor at a later date.
