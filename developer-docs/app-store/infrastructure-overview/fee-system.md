# Fee System Documentation

## TL;DR

The Vultisig fee system enables plugins to monetize through three fee models: per-transaction fees (charged for each operation), subscription fees (recurring monthly/yearly), and per-installation fees (one-time charges). Fees are collected either immediately or in batches, automatically converted to USDC via DEX aggregators, and distributed with 70% to developers and 30% to the \$VULT treasury. The system integrates with the Verifier for TSS-secured automation and maintains transparent revenue tracking.

## Overview

The Vultisig fee system is a comprehensive solution for collecting usage fees from app users and managing treasury operations. It supports multiple fee structures, automatic token conversions, and transparent revenue distribution while maintaining the security guarantees of the TSS architecture.

## Fee System Architecture

```
┌─────────────────┐    ┌─────────────────┐
│      Fees       │    │   Verifier      │
│                 │    │                 │
│ • Fee           │◄──►│ • Automation    │
│   calculation   │    │   validation    │
│ • Collection    │    │ • TSS signing   │
│ • Conversion    │    │                 │
│ • Distribution  │    └─────────────────┘
└─────────────────┘
        │
┌─────────────────┐
│  DEX Aggregator │
│                 │
│ • Token swaps   │
│ • Best prices   │
│ • Slippage      │
│   protection    │
└─────────────────┘
        │
┌─────────────────┐
│ Vultisig        │
│ Treasury        │
│                 │
│ • USDC storage  │
│ • Revenue       │
│   tracking      │
└─────────────────┘
```

## Fee Types and Structures

### 1. Per-Transaction Fees

**Overview:** Charged each time your app executes a transaction on behalf of users.

```json
{
  "fee_type": "per_transaction",
  "amount": "1000000",
  "denomination": "usdc",
  "collection_frequency": "immediate",
  "description": "Fee charged per executed transaction"
}
```

**Best For:** Trading bots, automation plugins, swap aggregators, portfolio rebalancers.

### 2. Subscription Fees

**Overview:** Fixed recurring charges at regular intervals, regardless of transaction count.

```json
{
  "fee_type": "subscription",
  "amount": "5000000",
  "denomination": "usdc",
  "collection_frequency": "monthly",
  "billing_period": "30d",
  "description": "Monthly subscription fee"
}
```
**Best For:** Portfolio analytics, premium features, unlimited access models, enterprise tools.

### 3. Per installation

**Overview:** One-time charge when users install your app from the App Store.

```json
{
  "fee_type": "per_transaction",
  "amount": "1000000",
  "denomination": "usdc",
  "collection_frequency": "immediate",
  "description": "Fee charged per installation"
}
```
**Best For:** Premium plugins, specialized tools, one-time setup services, license-based access.

## Fee Collection Mechanisms

- Immediate Collection (Fees are collected immediately when transactions are executed).
- Deferred Collection (Fees are accumulated and collected in batches)

## Token Conversion System

The fee system automatically converts various tokens to USDC for treasury management.
