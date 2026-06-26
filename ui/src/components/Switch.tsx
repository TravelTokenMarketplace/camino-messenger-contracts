import { type ReactNode } from "react";
import { Switch as HSwitch, Field, Label } from "@headlessui/react";

interface SwitchProps {
  checked: boolean;
  onChange: (checked: boolean) => void;
  disabled?: boolean;
  label?: ReactNode;
}

/** Styled toggle switch built on HeadlessUI, with an optional clickable label. */
export function Switch({ checked, onChange, disabled, label }: SwitchProps) {
  return (
    <Field
      disabled={disabled}
      className="flex items-center gap-2 text-sm data-[disabled]:cursor-not-allowed data-[disabled]:opacity-50"
    >
      {label && <Label className="cursor-pointer select-none">{label}</Label>}
      <HSwitch
        checked={checked}
        onChange={onChange}
        className="group relative inline-flex h-5 w-9 shrink-0 cursor-pointer rounded-full border border-tarmac-300 bg-tarmac-200 transition-colors focus:outline-none focus-visible:ring-2 focus-visible:ring-camino-500 data-[checked]:border-camino-600 data-[checked]:bg-camino-600 dark:border-tarmac-600 dark:bg-tarmac-700 dark:data-[checked]:border-camino-500 dark:data-[checked]:bg-camino-500"
      >
        <span className="pointer-events-none absolute left-0.5 top-0.5 inline-block h-3.5 w-3.5 transform rounded-full bg-paper-raised shadow transition-transform group-data-[checked]:translate-x-4" />
      </HSwitch>
    </Field>
  );
}
