import React from "react";
import { useLocation, useHistory } from "react-router-dom";
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import { Button } from "@material-ui/core";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      '& > *': {
        margin: theme.spacing(1),
        width: '25ch',
      },
    },
  }),
);

type Props = {
  id: number;
  owner_id: number;
  start_time: number;
  end_time: number;
  meeting_id: string;
  price: number;
};

const StudentLesson: React.FC<Props> = (props) => {
  const { id, owner_id, start_time, end_time, meeting_id, price } = props;
  const path = useLocation().pathname;
  const history = useHistory();
  const classes = useStyles();

  const start_time_jst = new Date(start_time * 1000).toLocaleString();
  const end_time_jst = new Date(end_time * 1000).toLocaleString();

  const clickHander = () => {
    const reserveFormPath = path + "/" + id;
    history.push(reserveFormPath);
  };

  const studentLessonList = {
    textAlign: 'center' as const,
  }

  return (
      <ul>
        <li>開始時間: {start_time_jst}</li>
        <li>終了時間: {end_time_jst}</li>
        <li>料金: {price}円</li>
        <p>
          <Button variant="contained" color="primary" onClick={clickHander}>予約する</Button>
        </p>
      </ul>
  );
};

export default StudentLesson;
