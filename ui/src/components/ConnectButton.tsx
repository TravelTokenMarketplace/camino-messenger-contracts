import { LogOut, Wallet } from "lucide-react";
import { useAccount, useConnect, useDisconnect } from "wagmi";
import { shortAddress } from "../lib/format";

export function ConnectButton() {
  const { address, isConnected } = useAccount();
  const { connect, connectors } = useConnect();
  const { disconnect } = useDisconnect();

  if (isConnected && address)
    return (
      <button
        className="inline-flex items-center gap-1.5 rounded border px-3 py-1.5 dark:border-gray-700"
        onClick={() => disconnect()}
      >
        <span className="font-mono text-sm">{shortAddress(address)}</span>
        <LogOut className="h-4 w-4 text-gray-400" />
      </button>
    );

  return (
    <button
      className="inline-flex items-center gap-1.5 rounded bg-indigo-600 px-3 py-1.5 text-white hover:bg-indigo-700"
      onClick={() => connect({ connector: connectors[0] })}
    >
      <Wallet className="h-4 w-4" /> Connect Wallet
    </button>
  );
}
