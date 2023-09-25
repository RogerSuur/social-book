import { useState } from "react";
import { Nav } from "react-bootstrap";
import SearchUtility from "./SearchUtility";
import GenericModal from "./GenericModal";
import SearchBar from "./SearchBar";
import SearchResults from "./SearchResults";

const SearchSmallUtility = () => {
  const [show, setShow] = useState(false);
  const [searchResults, setSearchResults] = useState([]);

  const handleShow = () => setShow(true);
  const handleClose = () => setShow(false);

  return (
    <div className="d-md-none">
      <Nav.Link>
        <GenericModal
          linkText={"Search"}
          headerButton={<SearchBar setSearchResults={setSearchResults} />}
        >
          <>
            {searchResults.length > 0 && (
              <SearchResults
                searchResults={searchResults}
                setSearchResults={setSearchResults}
              />
            )}
          </>
        </GenericModal>
      </Nav.Link>
    </div>
  );
};

export default SearchSmallUtility;
