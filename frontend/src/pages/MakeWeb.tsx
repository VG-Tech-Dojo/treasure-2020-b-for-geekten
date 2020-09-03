import React, { useState } from "react";
import useUser from "../hooks/useUser";
import { createWebsite } from "./../api";
import ErrorMessage from "./../components/ErrorMessage";
import firebase from "firebase/app";
import {zoomConfig} from "../../env";
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import { Button } from '@material-ui/core';
import TextField from '@material-ui/core/TextField';

interface State {
  title: string;
  profile: string;
  theme: string;
  content: string;
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

const MakeWeb: React.FC = () => {
  const [madeWeb, setMadeWeb] = useState(false);
  const [values, setValues] = useState<State>({
    title: "",
    profile: "",
    theme: "default",
    content: "",
  });
  const [errorMessage, setErrorMessage] = useState(null);
  const { user } = useUser();
  const classes = useStyles();

  const handleChange = (prop: keyof State) => (e: React.ChangeEvent<HTMLInputElement>) => {
    setValues({ ...values, [prop]: e.target.value });
  };

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const firebaseUser = firebase.auth().currentUser!;
    const token = await user!.getIdToken();
    const newWebsite = await createWebsite(token, values)
      .then(() => setMadeWeb(true))
      .catch((err) =>
        setErrorMessage(err.message)
    );
    console.log(newWebsite);
    console.log(firebaseUser.uid);
    window.location.href = `https://zoom.us/oauth/authorize?response_type=code&client_id=${zoomConfig.zoomClientID}&redirect_uri=${encodeURIComponent(zoomConfig.zoomRedirectURI)}/owner/${firebaseUser.uid}/zoom_auth`
  };

  const containerStyle = {
    justifyContent: "center" as const,
    display: "flex",
  }

  const innerStyle = {
    display: "flex",
    flexDirection: "column" as const,
  }

  return (
    <div style={containerStyle}>
      <form onSubmit={onSubmit}>
        <div style={innerStyle}>
          <h3>Webサイト作成</h3>
          <ErrorMessage message={errorMessage} />
          <TextField id="outlined-basic" label="タイトル" variant="outlined" onChange={handleChange("title")}/>
          <TextField id="outlined-basic" label="プロフィール" variant="outlined" multiline onChange={handleChange("profile")}/>
          <TextField id="outlined-basic" label="テーマ" variant="outlined" disabled defaultValue="default" />
          <TextField id="outlined-basic" label="内容" variant="outlined" multiline onChange={handleChange("content")}/>
          <p>
            <Button type="submit" variant="contained" color="primary">送信</Button>
          </p>
        </div>
      </form>
    </div>
  );
};

export default MakeWeb;
