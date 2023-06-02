import React, { useState, useEffect } from "react";
import axios from "axios";

const List = ({ url, mapFunction }) => {
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
  }, []);
  const renderedList = listData.map(mapFunction);
  //console.log(renderedList[0].key);
  return <div>{renderedList}</div>;
};
export default List;
