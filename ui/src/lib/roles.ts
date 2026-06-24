import { keccak256, toBytes, type Hex } from "viem";

export const ACCOUNT_ROLES = [
  "DEFAULT_ADMIN_ROLE",
  "UPGRADER_ROLE",
  "BOT_ADMIN_ROLE",
  "MESSENGER_BOT_ROLE",
  "GAS_WITHDRAWER_ROLE",
  "WITHDRAWER_ROLE",
  "BOOKING_OPERATOR_ROLE",
  "SERVICE_ADMIN_ROLE",
] as const;

export const MANAGER_ROLES = [
  "DEFAULT_ADMIN_ROLE",
  "PAUSER_ROLE",
  "UPGRADER_ROLE",
  "VERSIONER_ROLE",
  "SERVICE_REGISTRY_ADMIN_ROLE",
] as const;

export const BOOKINGTOKEN_ROLES = [
  "DEFAULT_ADMIN_ROLE",
  "UPGRADER_ROLE",
  "MIN_EXPIRATION_ADMIN_ROLE",
] as const;

export type RoleName =
  | (typeof ACCOUNT_ROLES)[number]
  | (typeof MANAGER_ROLES)[number]
  | (typeof BOOKINGTOKEN_ROLES)[number];

const ZERO_BYTES32 = `0x${"0".repeat(64)}` as Hex;

function compute(name: string): Hex {
  return name === "DEFAULT_ADMIN_ROLE" ? ZERO_BYTES32 : keccak256(toBytes(name));
}

export const ROLE_HASHES = Object.fromEntries(
  [...ACCOUNT_ROLES, ...MANAGER_ROLES, ...BOOKINGTOKEN_ROLES].map((r) => [r, compute(r)]),
) as Record<RoleName, Hex>;

export function roleHash(name: RoleName): Hex {
  return ROLE_HASHES[name];
}
