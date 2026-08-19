import { useEffect, useState, type KeyboardEvent } from 'react'
import type { Suggestion } from '../api/autocomplete'

export type UseComboboxNavigationResult = {
  highlightedIndex: number
  isOpen: boolean
  activeDescendantId: string | undefined
  getOptionId: (index: number) => string
  handleInputChange: (value: string) => void
  selectSuggestion: (value: string) => void
  handleKeyDown: (event: KeyboardEvent<HTMLInputElement>) => void
  setHighlightedIndex: (index: number) => void
}

const OPTION_ID_PREFIX = 'suggestion-option-'

export function useComboboxNavigation(
  suggestions: Suggestion[],
  onInputChange: (value: string) => void,
  onSuggestionSelect: (value: string) => void,
): UseComboboxNavigationResult {
  const [highlightedIndex, setHighlightedIndex] = useState(-1)
  const [isDismissed, setIsDismissed] = useState(false)

  useEffect(() => {
    setHighlightedIndex(-1)
  }, [suggestions])

  const isOpen = suggestions.length > 0 && !isDismissed

  const getOptionId = (index: number) => `${OPTION_ID_PREFIX}${index}`

  const activeDescendantId =
    isOpen && highlightedIndex >= 0
      ? getOptionId(highlightedIndex)
      : undefined

  const handleInputChange = (value: string) => {
    setIsDismissed(false)
    onInputChange(value)
  }

  const selectSuggestion = (value: string) => {
    setIsDismissed(true)
    onSuggestionSelect(value)
  }

  const handleKeyDown = (event: KeyboardEvent<HTMLInputElement>) => {
    if (!isOpen) {
      return
    }

    switch (event.key) {
      case 'ArrowDown':
        event.preventDefault()
        setHighlightedIndex((current) => (current + 1) % suggestions.length)
        break
      case 'ArrowUp':
        event.preventDefault()
        setHighlightedIndex(
          (current) => (current - 1 + suggestions.length) % suggestions.length,
        )
        break
      case 'Enter':
        if (highlightedIndex >= 0) {
          event.preventDefault()
          selectSuggestion(suggestions[highlightedIndex].value)
        }
        break
      case 'Escape':
        event.preventDefault()
        setIsDismissed(true)
        break
      default:
        break
    }
  }

  return {
    highlightedIndex,
    isOpen,
    activeDescendantId,
    getOptionId,
    handleInputChange,
    selectSuggestion,
    handleKeyDown,
    setHighlightedIndex,
  }
}
