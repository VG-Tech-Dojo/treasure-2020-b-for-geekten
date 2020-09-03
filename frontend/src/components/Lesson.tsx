import React from "react";
import {Paper, List, ListItem, Divider} from "@material-ui/core"

type Props = {
  id: number;
  owner_id: number;
  start_time: number;
  end_time: number;
  meeting_id: string;
  price: number;
};

const Lesson: React.FC<Props> = (props) => {
  const { start_time, end_time, meeting_id, price } = props;
  const start_time_jst = new Date(start_time * 1000).toLocaleString();
  const end_time_jst = new Date(end_time * 1000).toLocaleString();

  return (
    <List>
      <ListItem>[{start_time_jst}] 〜 [{end_time_jst}] (￥{price}円)</ListItem>
      <ListItem>ZoomミーティングURL:　<a href={meeting_id} target="_blank">{meeting_id}</a></ListItem>
    </List>
  );
};

export default Lesson;
