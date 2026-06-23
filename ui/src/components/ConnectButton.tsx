import { useAccount, useConnect, useDisconnect } from "wagmi";
import { shortAddress } from "../lib/format";

export function ConnectButton() {
  const { address, isConnected } = useAccount();
  const { connect, connectors } = useConnect();
  const { disconnect } = useDisconnect();

  if (isConnected && address)
    return (
      <button className="rounded border px-3 py-1.5" onClick={() => disconnect()}>
        {shortAddress(address)} · Disconnect
      </button>
    );

  return (
    <button
      className="rounded bg-indigo-600 px-3 py-1.5 text-white"
      onClick={() => connect({ connector: connectors[0] })}
    >
      Connect Wallet
    </button>
  );
}
