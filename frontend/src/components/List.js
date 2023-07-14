import React, { useState, useEffect } from "react";
import axios from "axios";

const List = ({ url, mapFunction, loadNewGroups }) => {
  const [listData, setListData] = useState([]);
  useEffect(() => {
    const fetchData = async () => {
      await axios
        .get(url, {
          withCredentials: true,
        })
        .then((response) => {
          setListData(response.data);
          console.log(url, response.data);
        });
    };
    fetchData();
  }, [url, loadNewGroups]);
  const renderedList = listData?.map(mapFunction);
  return <div>{renderedList}</div>;
};

export default List;
