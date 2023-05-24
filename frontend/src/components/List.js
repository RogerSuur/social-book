import React, { useState, useEffect } from "react";
import axios from "axios";

const List = ({ url, mapFunction }) => {
  const [listData, setListData] = useState([]);
  useEffect(() => {
    // console.log("list data request");
    const fetchData = async () => {
      await axios
        .get(url, {
          withCredentials: true,
        })
        .then((response) => setListData(response.data));
    };
    fetchData();
  }, []);
  const renderedList = listData.map(mapFunction);
  return <div>{renderedList}</div>;
};
export default List;
