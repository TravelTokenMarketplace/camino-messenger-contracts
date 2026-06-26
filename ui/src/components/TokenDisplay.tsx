import { shortAddress } from "../lib/format";
import { CopyButton } from "./CopyButton";
import { Identicon } from "./Identicon";
import { Tooltip } from "./Tooltip";

/**
 * Render an ERC20 token in a human-readable way on three stacked lines:
 * identicon + symbol headline (or compact address when the symbol is unknown),
 * the full name as muted secondary text, and the compacted address (full on
 * hover) with a copy button. Stacking keeps the symbol and name from being
 * truncated when the widget is in a narrow column.
 */
export function TokenDisplay({
  address,
  symbol,
  name,
  className = "",
}: {
  address: string;
  symbol?: string;
  name?: string;
  className?: string;
}) {
  const short = shortAddress(address);
  return (
    <span className={`inline-flex min-w-0 items-center gap-2 ${className}`}>
      <Identicon address={address} size={20} />
      <span className="flex min-w-0 flex-col">
        <span className="truncate font-medium leading-tight">{symbol ?? short}</span>
        {name && <span className="truncate text-xs leading-tight text-tarmac-500 dark:text-tarmac-400">{name}</span>}
        <span className="flex items-center gap-1">
          <Tooltip content={address}>
            <span className="font-mono text-[11px] leading-tight text-tarmac-500 dark:text-tarmac-400">{short}</span>
          </Tooltip>
          <CopyButton value={address} label="Copy address" />
        </span>
      </span>
    </span>
  );
}
