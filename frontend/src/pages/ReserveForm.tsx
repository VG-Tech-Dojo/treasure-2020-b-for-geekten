import React, { useEffect, useState } from "react";
import { useParams, useHistory } from "react-router-dom";
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import { Button } from '@material-ui/core';
import TextField from '@material-ui/core/TextField';

import { getLesson, reserveLesson } from "../api";
import ErrorMessage from "./../components/ErrorMessage";
import PayjpCheckout from "../components/PayjpCheckout";
import { payjpConfig } from "../../env";

interface State {
  first_name: string;
  last_name: string;
  email: string;
  paid_price: number;
  card_token: string;
}

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

const ReserveForm = () => {
  const [values, setValues] = useState<State>({
    first_name: "",
    last_name: "",
    email: "",
    paid_price: 0,
    card_token: "",
  });
  const [errorMessage, setErrorMessage] = useState(null);
  const [isLoading, setIsLoading] = useState(false);
  const handleChange = (prop: keyof State) => (e: React.ChangeEvent<HTMLInputElement>) => {
    setValues({ ...values, [prop]: e.target.value });
  };

  const { lesson_id } = useParams<{ lesson_id: string }>();
  const lessonID: string = lesson_id;
  const history = useHistory();
  const classes = useStyles();

  useEffect(() => {
    const fetchData = async () => {
      setIsLoading(true);

      const lesson = await getLesson(lessonID).catch((err) => {
        console.log(err);
        setErrorMessage(err.message);
      });
      if (lesson === "NotFound") {
        window.location.href = "/notfound";
      }
      setValues((prevValues) => {
        prevValues.paid_price = lesson.price;
        return prevValues;
      });

      setIsLoading(false);
    };
    fetchData().catch((err) => setErrorMessage(err.message));
  }, []);

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const newReservation = await reserveLesson(lessonID, values).catch((err) =>
      setErrorMessage(err.message)
    );
    console.log(newReservation);
    history.replace("/reserved");
  };

  function onCreateToken(response) {
    console.log("card token created");
    console.log("payjp response: ", response);
    // setValues(values)
    setValues((prevValues) => {
      prevValues.card_token = response.token;
      return prevValues;
    });
  }

  function onFailToken(statusCode, err) {
    console.log("failed to create card token");
    console.log("status code: ", statusCode);
    console.log("error: ", err);
    setErrorMessage(err.message);
  }

  const payjpCheckoutProps = {
    dataKey: payjpConfig.payjpPkTest,
    dataText: "カード情報を入力",
    dataPartial: "true",
    onCreatedHandler: onCreateToken,
    onFailHandler: onFailToken,
  };

  const containerStyle = {
    textAlign: "center" as const
  }

  return (
    <div style={containerStyle}>
      <ErrorMessage message={errorMessage} />
      <h3>生徒予約画面</h3>
      <form onSubmit={onSubmit}>
        <TextField id="outlined-basic" label="姓" variant="outlined" onChange={handleChange("last_name")}/>
        <TextField id="outlined-basic" label="名" variant="outlined" onChange={handleChange("first_name")}/>
        <TextField id="outlined-basic" label="メールアドレス" variant="outlined" onChange={handleChange("email")}/>
        {isLoading ? <div>Loading ...</div> : <div>￥{values.paid_price}</div>}
        <div>
          <PayjpCheckout {...payjpCheckoutProps} />
        </div>
        <p>
          <Button type="submit" variant="contained" color="primary">予約確定</Button>
        </p>
      </form>
    </div>
  );
};

export default ReserveForm;
