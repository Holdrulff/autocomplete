import { EmptyState } from './components/EmptyState'
import { SearchInput } from './components/SearchInput'
import { StatusMessage } from './components/StatusMessage'
import { SuggestionList } from './components/SuggestionList'
import { ThemeToggle } from './components/ThemeToggle'
import { useAutocomplete } from './hooks/useAutocomplete'
import { useComboboxNavigation } from './hooks/useComboboxNavigation'
import './App.css'

const LISTBOX_ID = 'suggestion-listbox'
const STATUS_ID = 'search-help'

function App() {
  const {
    prefix,
    setPrefix,
    selectPrefix,
    isSearchEnabled,
    normalizedPrefix,
    suggestions,
    isLoading,
    error,
  } = useAutocomplete()

  const {
    highlightedIndex,
    isOpen,
    activeDescendantId,
    getOptionId,
    handleInputChange,
    selectSuggestion,
    handleKeyDown,
    setHighlightedIndex,
  } = useComboboxNavigation(suggestions, setPrefix, selectPrefix)

  const showEmptyState =
    isSearchEnabled &&
    normalizedPrefix !== '' &&
    !isLoading &&
    !error &&
    suggestions.length === 0
  const showLoadingPopup =
    isSearchEnabled && isLoading && suggestions.length === 0
  const isPopupVisible = isOpen || showLoadingPopup

  return (
    <main className="app-shell">
      <section className="search-card" aria-labelledby="search-title">
        <div className="search-header">
          <div>
            <p className="eyebrow">Autocomplete</p>
            <h1 id="search-title">Find a technology</h1>
          </div>

          <ThemeToggle />
        </div>

        <p className="search-description">
          Start typing to explore popular technology tags from Stack Overflow.
        </p>

        <SearchInput
          value={prefix}
          onChange={handleInputChange}
          onKeyDown={handleKeyDown}
          listboxId={LISTBOX_ID}
          activeDescendantId={activeDescendantId}
          isExpanded={isPopupVisible}
          describedById={STATUS_ID}
        />

        <StatusMessage
          id={STATUS_ID}
          normalizedPrefix={normalizedPrefix}
          isLoading={isLoading}
          error={error}
          resultCount={suggestions.length}
        />

        {isPopupVisible && (
          <SuggestionList
            suggestions={suggestions}
            isLoading={isLoading}
            highlightedIndex={highlightedIndex}
            listboxId={LISTBOX_ID}
            getOptionId={getOptionId}
            onSelect={selectSuggestion}
            onHighlight={setHighlightedIndex}
          />
        )}

        {showEmptyState && <EmptyState normalizedPrefix={normalizedPrefix} />}
      </section>
    </main>
  )
}

export default App
