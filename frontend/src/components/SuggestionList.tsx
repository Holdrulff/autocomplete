import type { Suggestion } from '../api/autocomplete'
import { SuggestionItem } from './SuggestionItem'
import './SuggestionList.css'

const SKELETON_ROW_COUNT = 4

type SuggestionListProps = {
  suggestions: Suggestion[]
  isLoading: boolean
  highlightedIndex: number
  listboxId: string
  getOptionId: (index: number) => string
  onSelect: (value: string) => void
  onHighlight: (index: number) => void
}

export function SuggestionList({
  suggestions,
  isLoading,
  highlightedIndex,
  listboxId,
  getOptionId,
  onSelect,
  onHighlight,
}: SuggestionListProps) {
  if (isLoading && suggestions.length === 0) {
    return (
      <ul className="suggestion-list" id={listboxId} role="listbox" aria-busy="true">
        {Array.from({ length: SKELETON_ROW_COUNT }, (_, index) => (
          <li key={index} className="suggestion-skeleton-row" aria-hidden="true">
            <span className="skeleton-block skeleton-block-value" />
            <span className="skeleton-block skeleton-block-score" />
          </li>
        ))}
      </ul>
    )
  }

  if (suggestions.length === 0) {
    return null
  }

  return (
    <ul
      className="suggestion-list"
      id={listboxId}
      role="listbox"
      aria-label="Autocomplete suggestions"
    >
      {suggestions.map((suggestion, index) => (
        <SuggestionItem
          key={suggestion.value}
          suggestion={suggestion}
          index={index}
          id={getOptionId(index)}
          isHighlighted={index === highlightedIndex}
          onSelect={onSelect}
          onHighlight={onHighlight}
        />
      ))}
    </ul>
  )
}
