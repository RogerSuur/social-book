export const SearchResults = ({ searchResults }) => {
  return (
    <div className="results-list">
      {searchResults &&
        searchResults.map((result) => (
          <div>
            <a
              key={result.id}
              href={
                result.userId === 0
                  ? `/groups/${result.groupId}`
                  : `profile/${result.userId}`
              }
            >
              {result.name}
            </a>
          </div>
        ))}
    </div>
  );
};
