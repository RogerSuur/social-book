export const SearchResults = ({ searchResults }) => {
  return (
    <div className="results-list">
      {searchResults &&
        searchResults.map((result) => <div key={result.id}>{result.name}</div>)}
    </div>
  );
};
