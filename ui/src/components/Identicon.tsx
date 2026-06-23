import Jazzicon, { jsNumberForAddress } from "react-jazzicon";
import { isAddress } from "viem";

/**
 * Deterministic avatar for an address (the colorful "jazzicon" used by MetaMask),
 * making it easy to visually distinguish addresses at a glance.
 */
export function Identicon({ address, size = 16 }: { address: string; size?: number }) {
  if (!isAddress(address)) return null;
  return (
    <span className="inline-flex shrink-0 overflow-hidden rounded-full" style={{ width: size, height: size }}>
      <Jazzicon diameter={size} seed={jsNumberForAddress(address)} />
    </span>
  );
}
