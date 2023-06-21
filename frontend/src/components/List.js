import React, { useState, useEffect } from "react";
import axios from "axios";

const List = ({ url, mapFunction }) => {
  console.log(url);
  const [listData, setListData] = useState([]);
  useEffect(() => {
    const fetchData = async () => {
      await axios
        .get(url, {
          withCredentials: true,
        })
        .then((response) => setListData(response.data));
    };
    fetchData();
  }, [url]);
  const renderedList = listData.map(mapFunction);
  return <div>{renderedList}</div>;
};

export default List;
