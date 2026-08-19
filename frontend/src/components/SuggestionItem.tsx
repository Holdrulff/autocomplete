import type { CSSProperties } from 'react'
import type { Suggestion } from '../api/autocomplete'

const scoreFormatter = new Intl.NumberFormat('en-US')

type SuggestionItemProps = {
  suggestion: Suggestion
  index: number
  id: string
  isHighlighted: boolean
  onSelect: (value: string) => void
  onHighlight: (index: number) => void
}

export function SuggestionItem({
  suggestion,
  index,
  id,
  isHighlighted,
  onSelect,
  onHighlight,
}: SuggestionItemProps) {
  return (
    <li
      role="option"
      id={id}
      aria-selected={isHighlighted}
      className={
        isHighlighted ? 'suggestion-item is-highlighted' : 'suggestion-item'
      }
      style={{ '--stagger-index': index } as CSSProperties}
      onMouseDown={(event) => {
        event.preventDefault()
      }}
      onClick={() => {
        onSelect(suggestion.value)
      }}
      onMouseEnter={() => onHighlight(index)}
    >
      <span className="suggestion-value">{suggestion.value}</span>
      <span className="suggestion-score">
        {scoreFormatter.format(suggestion.score)} questions
      </span>
    </li>
  )
}
