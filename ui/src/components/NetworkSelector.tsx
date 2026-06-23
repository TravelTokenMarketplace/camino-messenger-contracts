import { Listbox, ListboxButton, ListboxOption, ListboxOptions } from "@headlessui/react";
import { AlertTriangle, Check, ChevronsUpDown } from "lucide-react";
import { useAccount, useSwitchChain } from "wagmi";
import { ENABLED_CHAINS } from "../config/chains";
import { useActiveChain } from "../wallet/activeChain";

export function NetworkSelector() {
  const { activeChainId, setActiveChainId } = useActiveChain();
  const { isConnected, chainId: walletChainId } = useAccount();
  const { switchChain } = useSwitchChain();

  function onChange(id: number) {
    if (isConnected) switchChain({ chainId: id });
    else setActiveChainId(id);
  }

  const walletUnsupported =
    isConnected && walletChainId !== undefined && !ENABLED_CHAINS.some((c) => c.id === walletChainId);
  const active = ENABLED_CHAINS.find((c) => c.id === activeChainId);

  return (
    <Listbox value={activeChainId} onChange={onChange}>
      <div className="relative">
        <ListboxButton className="inline-flex items-center gap-2 rounded-md border border-gray-300 bg-white px-3 py-1.5 text-sm transition-colors hover:bg-gray-50 dark:border-gray-700 dark:bg-gray-800 dark:text-gray-100 dark:hover:bg-gray-700">
          {walletUnsupported ? (
            <AlertTriangle className="h-4 w-4 text-red-500" />
          ) : (
            <span className="h-2 w-2 rounded-full bg-emerald-500" aria-hidden />
          )}
          <span>{walletUnsupported ? "Unsupported network" : active?.name ?? "Network"}</span>
          <ChevronsUpDown className="h-4 w-4 text-gray-400" />
        </ListboxButton>
        <ListboxOptions
          anchor="bottom end"
          className="z-30 mt-1 w-52 rounded-md border border-gray-200 bg-white py-1 text-sm shadow-lg focus:outline-none dark:border-gray-700 dark:bg-gray-800"
        >
          {walletUnsupported && (
            <div className="flex items-start gap-2 px-3 py-2 text-xs text-red-600 dark:text-red-400">
              <AlertTriangle className="mt-0.5 h-3.5 w-3.5 shrink-0" />
              Your wallet is on an unsupported network. Pick one below to switch.
            </div>
          )}
          {ENABLED_CHAINS.map((c) => (
            <ListboxOption
              key={c.id}
              value={c.id}
              className="flex cursor-pointer items-center justify-between gap-2 px-3 py-1.5 text-gray-700 data-[focus]:bg-gray-100 dark:text-gray-200 dark:data-[focus]:bg-gray-700"
            >
              <span>{c.name}</span>
              {c.id === activeChainId && <Check className="h-4 w-4 text-indigo-500" />}
            </ListboxOption>
          ))}
        </ListboxOptions>
      </div>
    </Listbox>
  );
}
