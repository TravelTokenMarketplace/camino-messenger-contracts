import { type ReactNode, useState } from "react";
import { Plus, Trash2 } from "lucide-react";
import { Autocomplete } from "./Autocomplete";
import { Card } from "./Card";
import { Input } from "./Input";
import { RoleGate } from "./RoleGate";
import { RowAction } from "./RowAction";
import { TxButton } from "./TxButton";

interface ListManagerProps {
  title: string;
  items: string[];
  isLoading: boolean;
  roleName: string;
  hasRole: boolean;
  addLabel: string;
  addPlaceholder: string;
  onAdd: (value: string) => Promise<`0x${string}`>;
  onRemove: (value: string) => Promise<`0x${string}`>;
  onChanged?: () => void;
  renderItem?: (value: string) => ReactNode;
  /** Optional autocomplete suggestions for the add input. */
  suggestions?: string[];
}

export function ListManager(props: ListManagerProps) {
  const { title, items, isLoading, roleName, hasRole, addLabel, addPlaceholder, suggestions } = props;
  const [value, setValue] = useState("");

  return (
    <Card title={title}>
      {isLoading ? (
        <p>Loading…</p>
      ) : (
        <ul className="mb-4 divide-y">
          {items.length === 0 && <li className="py-2 text-sm text-tarmac-400">None</li>}
          {items.map((item) => (
            <li key={item} className="group flex items-center justify-between gap-3 py-2">
              <span className="min-w-0 font-mono text-sm">{props.renderItem ? props.renderItem(item) : item}</span>
              {hasRole && (
                <RowAction>
                  <TxButton
                    label="Remove"
                    variant="danger"
                    icon={<Trash2 className="h-4 w-4" />}
                    write={() => props.onRemove(item)}
                    onConfirmed={props.onChanged}
                  />
                </RowAction>
              )}
            </li>
          ))}
        </ul>
      )}
      <RoleGate hasRole={hasRole} roleName={roleName} action={addLabel}>
        <div className="flex items-end gap-2">
          {suggestions ? (
            <Autocomplete
              className="flex-1"
              value={value}
              onChange={setValue}
              options={suggestions}
              placeholder={addPlaceholder}
            />
          ) : (
            <Input
              className="flex-1 font-mono"
              placeholder={addPlaceholder}
              value={value}
              onChange={(e) => setValue(e.target.value)}
            />
          )}
          <TxButton
            label={addLabel}
            icon={<Plus className="h-4 w-4" />}
            disabled={!value}
            write={() => props.onAdd(value)}
            onConfirmed={() => {
              setValue("");
              props.onChanged?.();
            }}
          />
        </div>
      </RoleGate>
    </Card>
  );
}
