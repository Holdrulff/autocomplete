import type { ChangeEvent, KeyboardEvent } from 'react'
import { ClearIcon, SearchIcon } from './icons'
import './SearchInput.css'

type SearchInputProps = {
  value: string
  onChange: (value: string) => void
  onKeyDown: (event: KeyboardEvent<HTMLInputElement>) => void
  listboxId: string
  activeDescendantId: string | undefined
  isExpanded: boolean
  describedById: string
}

export function SearchInput({
  value,
  onChange,
  onKeyDown,
  listboxId,
  activeDescendantId,
  isExpanded,
  describedById,
}: SearchInputProps) {
  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    onChange(event.target.value)
  }

  return (
    <div className="search-field">
      <label htmlFor="autocomplete-input">Technology</label>

      <div className="search-combobox-wrapper">
        <SearchIcon className="search-field-icon" />

        <input
          id="autocomplete-input"
          name="autocomplete"
          type="text"
          role="combobox"
          placeholder="e.g., java, react, python"
          autoComplete="off"
          value={value}
          onChange={handleChange}
          onKeyDown={onKeyDown}
          aria-describedby={describedById}
          aria-autocomplete="list"
          aria-expanded={isExpanded}
          aria-controls={listboxId}
          aria-activedescendant={activeDescendantId}
        />

        {value !== '' && (
          <button
            type="button"
            className="search-clear-btn"
            aria-label="Clear search"
            onClick={() => onChange('')}
          >
            <ClearIcon />
          </button>
        )}
      </div>
    </div>
  )
}
