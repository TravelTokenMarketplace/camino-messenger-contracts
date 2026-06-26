import { useMemo, useState } from "react";
import { Combobox, ComboboxButton, ComboboxInput, ComboboxOption, ComboboxOptions } from "@headlessui/react";
import { Check, ChevronsUpDown, Search } from "lucide-react";
import Fuse from "fuse.js";

interface AutocompleteProps {
  value: string;
  onChange: (value: string) => void;
  options: string[];
  placeholder?: string;
  id?: string;
  className?: string;
  /** Allow submitting a value that isn't in the option list (default: true). */
  allowCustom?: boolean;
}

const inputClass =
  "w-full rounded-md border border-tarmac-300 bg-paper-raised py-1.5 pl-3 pr-16 text-sm text-tarmac-900 placeholder-tarmac-400 transition-colors focus:border-camino-500 focus:outline-none focus:ring-1 focus:ring-camino-500 dark:border-tarmac-600 dark:bg-tarmac-900 dark:text-tarmac-100 dark:placeholder-tarmac-500";

/**
 * Styled, fuzzy-searchable combobox. Opens on focus to reveal the available
 * options, filters as you type (typo-tolerant via Fuse), and — unlike a native
 * <datalist> — has a bounded, scrollable, themed dropdown. Free text is allowed
 * by default so values not in the list can still be entered.
 */
export function Autocomplete({
  value,
  onChange,
  options,
  placeholder,
  id,
  className,
  allowCustom = true,
}: AutocompleteProps) {
  const [query, setQuery] = useState("");
  const fuse = useMemo(() => new Fuse(options, { threshold: 0.4, ignoreLocation: true }), [options]);

  const q = query.trim();
  const filtered = (q ? fuse.search(q).map((r) => r.item) : options).slice(0, 50);

  return (
    <Combobox
      immediate
      value={value}
      onChange={(v: string | null) => {
        onChange(v ?? "");
        setQuery("");
      }}
    >
      <div className={`relative ${className ?? ""}`}>
        <ComboboxInput
          id={id}
          className={inputClass}
          placeholder={placeholder}
          displayValue={(v: string) => v}
          onChange={(e) => {
            setQuery(e.target.value);
            if (allowCustom) onChange(e.target.value);
          }}
        />
        <ComboboxButton className="absolute inset-y-0 right-0 flex items-center px-2 text-tarmac-400 hover:text-tarmac-600 dark:hover:text-tarmac-200">
          <ChevronsUpDown className="h-4 w-4" />
        </ComboboxButton>

        <ComboboxOptions className="absolute z-30 mt-1 max-h-60 w-full overflow-auto rounded-md border border-tarmac-200 bg-paper-raised py-1 text-sm shadow-lg focus:outline-none dark:border-tarmac-700 dark:bg-tarmac-800">
          {options.length === 0 ? (
            <div className="px-3 py-2 text-tarmac-400">No registered services found on this network.</div>
          ) : filtered.length === 0 ? (
            <div className="flex items-center gap-2 px-3 py-2 text-tarmac-400">
              <Search className="h-3.5 w-3.5" /> No match{q && ` for “${q}”`}.
            </div>
          ) : (
            filtered.map((opt) => (
              <ComboboxOption
                key={opt}
                value={opt}
                className="flex cursor-pointer items-center justify-between gap-2 px-3 py-1.5 text-tarmac-700 data-[focus]:bg-camino-600 data-[focus]:text-white dark:text-tarmac-200"
              >
                <span className="truncate font-mono">{opt}</span>
                <Check className="h-4 w-4 shrink-0 opacity-0 group-data-[selected]:opacity-100 data-[selected]:opacity-100" />
              </ComboboxOption>
            ))
          )}
        </ComboboxOptions>
      </div>
    </Combobox>
  );
}
