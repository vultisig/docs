---
description: >-
  Vultisig Airdrop V2: Enhanced rewards system. Earn $VULT based on vault
  value × time held. Achievement tiers, daily drops, and multipliers.
cover: ../../.gitbook/assets/Vultisig - 1200x630 - 56- 2x.png
coverY: 0
---

# Airdrop V2

Early adopters of the Vultisig wallet security standard are highly valued by the project, so Vultisig wants to give back to these high conviction members.&#x20;

The future $VULT airdrop will be proportional:

$$
vault\_asset\_value*time\_in\_vault
$$

While taking into account other multipliers, such as swap volume and referral multipliers, to further increase the virality.\
See further calculation [here](airdrop-v2.md#calculations).

***

## Airdrop Process

To register their vaults for airdrop, users send their exported vault public keys to an airdrop registry, which checks for the presence of funds on chains and then begins counting the airdrop value.&#x20;

This will include most actions performed in Vultisig (such as binding to nodes and staking, etc.).&#x20;

**A new scan is performed every cycle (day), and the user's airdrop value accumulates in the form of VULTIES (airdrop points).**&#x20;

Therefore, users who keep their funds in Vultisig the longest and use the vault actively will earn the highest airdrop value.&#x20;

The airdrop points are reset after each [season](airdrop-v2.md#seasons) and can be multiplied with special multipliers, which are listed [below](airdrop-v2.md#multipliers). The final share of the airdrop per season is the pro rata share of the airdrop allocation.

***

## Eligible assets

Assets must have a dollar value from an external provider in order to accumulate VULTIES (airdrop points).

<figure><img src="../../.gitbook/assets/image (15).png" alt=""><figcaption><p>BTC</p></figcaption></figure>

Supported Assets are:

* L1 tokens supported by Vultisig
* Tokens supported by 1inch, Jupiter and displayed on the airdrop website
* THORChain Bonds, LP(dual and single)
* MAYA Protocol Bonds, LP(dual and single)
* vTHOR
* Staked TON

Assets will be auto-discovered. In the case of non-discovery, assets can be added manually to verify if they are eligible.

***

## Airdrop amount

The airdrop will have a total amount of **6,000,000 $VULT** (6% of total supply) and will be split into:

Season 0 : 1,000,000 $VULT\
Season 1-4: 5,000,000 $VULT, distributed at the end of each season.

***

## Seasons

In order to create a more dynamic and inclusive airdrop system, a seasonal structure will be implemented, giving both early supporters and newcomers meaningful opportunities to participate.&#x20;

Each season will last 1 quarter and will have an airdrop share of 1,250,000 $VULT to distribute.&#x20;

**Airdrop points will be reset after each season.**

***

## Multipliers

To make the airdrop more interesting, several multipliers will be introduced starting with Season 1.

### **Swap Volume Multiplier**

To incentivize trading and make the airdrop more dynamic, a swap multiplier will be introduced. This will reward active users of Vultisig with a higher airdrop share based on their swap volume. The multiplier will reset each season to prevent excessive advantages and maintain fairness.

### **Referral Multiplier**

The [Referral Program](airdrop-v2.md#referral-program) is tracking referred wallets via the [Referral Telegram bot](https://docs.vultisig.com/other/vultisig-bot) of registered users.&#x20;

For a referral to be considered valid, the referred wallet must fund with at least $50 and take part in the airdrop.

{% hint style="info" %}
The referred wallet must continuously have at least $50 in the vault; otherwise, the referral will be voided.
{% endhint %}

This multiplier logarithmically multiplies and caps at 500 referrals, providing up to a maximum of 2x boost in your total airdrop points. The referral bonus, in contrast to other multipliers, will not reset between seasons, so you can build and retain your network advantage between seasons.

### Dedicated tokens and community multipliers

To attract more attention and expand the Vultisquad community, a special multiplier will be introduced for holding selected tokens or NFTs that will be selected in the coming future.&#x20;

The first token to receive this treatment will be $VULT itself, which will grant holders a 1.5x multiplier on its dollar value within the airdrop calculation.

The following tokens and NFTs have a higher multiplier:

| Token  | Multiplier |
| ------ | ---------- |
| $VULT  | 1.5        |
| $RUNE  | 1.3        |

***

| NFT        | Multiplier |
| ---------- | ---------- |
| Thorguards | 1.3        |



***

## Calculations

### Daily airdrop points per user

$$
user\_VULTIES += SQRT(\$\text{total\_vault\_value})
$$

### Swap Volume multiplier

$$
\text{swap\_volume\_multiplier} = 1+0.002*SQRT(\text{swap\_volume\_per\_season})
$$

### Referral multiplier

$$
\text{referral\_multiplier} = \min\left(2, 1+\frac{\log(1+\text referred\_wallets)}{\log(1+500)}\right)
$$

### Adjusted airdrop points per user

$$
\text adjusted\_user\_VULTIES = \text user\_VULTIES\times \text {swap\_volume\_multiplier} \times \text{referral\_multiplier}
$$

### Seasonal $VULT share

$$
\text{user\_airdrop\_share} = 1,250,000 \times \left( \frac{\text{adjusted\_user\_VULTIES}}{\text{total\_adjusted\_VULTIES}} \right)
$$

***

## How to register for the airdrop

* Download and open your Vultisig app

<figure><img src="../../.gitbook/assets/Frame 1000005130.png" alt="" width="279"><figcaption></figcaption></figure>

* [Export](../../vultisig-app-actions/managing-your-vault/vault-qr.md) the Vault QR of the Vultisig Vault

<figure><img src="../../.gitbook/assets/Frame 1000005131.png" alt="" width="279"><figcaption></figcaption></figure>

*   Connect to the [airdrop](https://airdrop.vultisig.com/import) page with your [Vultisig Extension](https://chromewebstore.google.com/detail/ggafhcdaplkhmmnlbfjpnnkepdfjaelb?utm_source=item-share-cp) or with uploading your Vault QR\


    <figure><img src="../../.gitbook/assets/image (11).png" alt="" width="375"><figcaption></figcaption></figure>
*   Join the Airdrop with your connected Vault with clicking `Join Airdrop` on the [web app](../../vultisig-ecosystem/web-app.md)\


    <figure><img src="../../.gitbook/assets/Button.png" alt="" width="225"><figcaption></figcaption></figure>
* Go to the `Balances` tab and check for auto-discovery or enable **ALL a**ssets you want to have counted towards the airdrop
* Earn VULTIES (Airdrop points) and track the leaderboard
* Register multiple Vaults

{% hint style="warning" %}
You need to enable all Chains and Tokens you want to have counted towards the airdrop once!
{% endhint %}

***

### Referral Program

Vultisig team has built a [Telegram bot](../../other/vultisig-bot.md) to easily invite friends and track the referred people.

* Invite people
* Let them install Vultisig
* Funding the wallet with at least $50 of assets
* **Increase your referral count and airdrop share**

**The referred people will have a strong impact on the multiplier.**

## Privacy

{% hint style="warning" %}
Do not register for the Airdrop if you do not wish for your public keys to be sent to the Airdrop Registry.\
**Note, after the Airdrop has finished, your public keys will be purged. No other user information is collected.**
{% endhint %}

Read the [#privacy-policy](../../other/privacy.md#privacy-policy "mention")
