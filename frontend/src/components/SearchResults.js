import { Link } from "react-router-dom";
export const SearchResults = ({ searchResults }) => {
  const searchResultsMap = searchResults.map((result, index) => (
    <div key={index}>
      <Link
        to={
          result.userId === 0
            ? `/groups/${result.groupId}`
            : `/profile/${result.userId}`
        }
      >
        {result.name}
      </Link>
    </div>
  ));

  return (
    <div className="results-list">
      <ul>{searchResultsMap}</ul>
    </div>
  );
};
