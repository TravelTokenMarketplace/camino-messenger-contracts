import { useState } from "react";
import { Menu, MenuButton, MenuItem, MenuItems } from "@headlessui/react";
import { Check, ChevronDown, Copy, ExternalLink, LogOut, Wallet } from "lucide-react";
import { useAccount, useConnect, useDisconnect } from "wagmi";
import { APP_CHAINS } from "../config/chains";
import { explorerAddrUrl, shortAddress } from "../lib/format";
import { Identicon } from "./Identicon";

export function ConnectButton() {
  const { address, isConnected, chainId } = useAccount();
  const { connect, connectors } = useConnect();
  const { disconnect } = useDisconnect();
  const [copied, setCopied] = useState(false);

  if (!isConnected || !address)
    return (
      <button
        className="inline-flex items-center gap-1.5 rounded-md bg-camino-600 px-3 py-1.5 text-white transition-colors hover:bg-camino-700"
        onClick={() => connect({ connector: connectors[0] })}
      >
        <Wallet className="h-4 w-4" /> Connect Wallet
      </button>
    );

  const explorer = APP_CHAINS.find((c) => c.id === chainId)?.explorerUrl;

  async function copy() {
    try {
      await navigator.clipboard?.writeText(address!);
      setCopied(true);
      setTimeout(() => setCopied(false), 1200);
    } catch {
      // clipboard unavailable; ignore
    }
  }

  return (
    <Menu as="div" className="relative">
      <MenuButton className="inline-flex items-center gap-1.5 rounded-md border border-tarmac-300 px-3 py-1.5 transition-colors hover:bg-tarmac-50 dark:border-tarmac-700 dark:hover:bg-tarmac-800">
        <Identicon address={address} />
        <span className="font-mono text-sm">{shortAddress(address)}</span>
        <ChevronDown className="h-4 w-4 text-tarmac-400" />
      </MenuButton>
      <MenuItems
        anchor="bottom end"
        className="z-30 mt-1 w-48 rounded-md border border-tarmac-200 bg-paper-raised py-1 text-sm shadow-lg focus:outline-none dark:border-tarmac-700 dark:bg-tarmac-800"
      >
        <MenuItem>
          <button
            onClick={copy}
            className="flex w-full items-center gap-2 px-3 py-1.5 text-left text-tarmac-700 data-[focus]:bg-tarmac-100 dark:text-tarmac-200 dark:data-[focus]:bg-tarmac-700"
          >
            {copied ? <Check className="h-4 w-4 text-green-600" /> : <Copy className="h-4 w-4 text-tarmac-400" />}
            {copied ? "Copied" : "Copy address"}
          </button>
        </MenuItem>
        {explorer && (
          <MenuItem>
            <a
              href={explorerAddrUrl(explorer, address)}
              target="_blank"
              rel="noreferrer"
              className="flex w-full items-center gap-2 px-3 py-1.5 text-left text-tarmac-700 data-[focus]:bg-tarmac-100 dark:text-tarmac-200 dark:data-[focus]:bg-tarmac-700"
            >
              <ExternalLink className="h-4 w-4 text-tarmac-400" /> View on explorer
            </a>
          </MenuItem>
        )}
        <div className="my-1 border-t border-tarmac-100 dark:border-tarmac-700" />
        <MenuItem>
          <button
            onClick={() => disconnect()}
            className="flex w-full items-center gap-2 px-3 py-1.5 text-left text-red-600 data-[focus]:bg-red-50 dark:text-red-400 dark:data-[focus]:bg-red-950"
          >
            <LogOut className="h-4 w-4" /> Disconnect
          </button>
        </MenuItem>
      </MenuItems>
    </Menu>
  );
}
