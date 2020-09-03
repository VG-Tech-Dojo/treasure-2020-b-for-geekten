import React, { useState, useEffect } from "react";
import { Button } from '@material-ui/core';
import ErrorMessage from "../components/ErrorMessage";
import useUser from "../hooks/useUser";
import { createOwner } from "../api";
import { useHistory } from "react-router-dom";
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';

interface State {
  first_name: string;
  last_name: string;
  postal_number: string;
  prefecture: string;
  city: string;
  address: string;
  address_optional: string;
  phone_number: string;
  email: string;
  bank_account_number: string;
  bank_branch_code: string;
  bank_code: string;
  bank_account_holder_name: string;
  bank_account_type: string;
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

const AccountForm: React.FC = () => {
  const { user } = useUser();
  const classes = useStyles();
  const [createdAccount, setCreatedAccount] = useState(false);
  const [errorMessage, setErrorMessage] = useState(null);
  const history = useHistory();
  const [values, setValues] = useState<State>({
    first_name: "",
    last_name: "",
    postal_number: "",
    prefecture: "",
    city: "",
    address: "",
    address_optional: "",
    phone_number: "",
    email: "",
    bank_account_number: "",
    bank_branch_code: "",
    bank_code: "",
    bank_account_holder_name: "",
    bank_account_type: "",
  });

  useEffect(() => {
    if (user === null) {
      return;
    }
    setValues({ ...values, ["email"]: user.email! });
  }, [user]);

  const handleChange = (prop: keyof State) => (e: React.ChangeEvent<HTMLInputElement>) => {
    setValues({ ...values, [prop]: e.target.value });
  };

  const onSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const token = await user!.getIdToken();
    const owner = await createOwner(token, values)
      .then(() => {
        history.replace('/make_web');
      })
      .then(() => setCreatedAccount(true))
      .catch((err) => {
        setErrorMessage(err.message);
      });
  };

  const formStyle = {
    display: "flex",
    flexDirection: "column" as const,
  }

  const containerStyle = {
    display: "flex",
    height: "120vh",
    justifyContent: "center",
    backgroundColor: "#FFFDF3",
  }

  const ownerInfoStyle = {
    display: "flex",
    flexDirection: "column" as const,
  }

  const bankInfoStyle = {
    display: "flex",
    textAlign: "center" as const,
    flexDirection: "column" as const,
  }

  return (
    <div style={containerStyle}>
      <form style={formStyle} onSubmit={onSubmit} className={classes.root} noValidate autoComplete="off">
        <div style={ownerInfoStyle}>
          <h1>講師登録</h1>
          <ErrorMessage message={errorMessage} />
          <TextField id="outlined-basic" label="姓" variant="outlined" onChange={handleChange("last_name")}/>
          <TextField id="outlined-basic" label="名" variant="outlined" onChange={handleChange("first_name")}/>
          <TextField id="outlined-basic" placeholder="1234567" label="郵便番号" variant="outlined" onChange={handleChange("postal_number")}/>
          <TextField id="outlined-basic" label="都道府県" variant="outlined" onChange={handleChange("prefecture")}/>
          <TextField id="outlined-basic" label="市区町村" variant="outlined" onChange={handleChange("city")}/>
          <TextField id="standard-full-width" fullWidth label="番地" variant="outlined" onChange={handleChange("address")}/>
          <TextField id="standard-full-width" fullWidth label="マンション名・部屋番号等" variant="outlined" onChange={handleChange("address_optional")}/>
          <TextField id="outlined-basic" placeholder="08011112222" label="電話番号" variant="outlined" onChange={handleChange("phone_number")}/>
        </div>
        <div style={bankInfoStyle}>
            <div>
            <TextField id="outlined-basic" label="口座登録氏名" variant="outlined" onChange={handleChange("bank_account_holder_name")}/>
            <TextField id="outlined-basic" label="銀行コード" variant="outlined" onChange={handleChange("bank_code")}/>
            <TextField id="outlined-basic" label="支店コード" variant="outlined" onChange={handleChange("bank_branch_code")}/>
            <TextField id="outlined-basic" label="口座番号" variant="outlined" onChange={handleChange("bank_account_number")}/>
            <TextField id="outlined-basic" label="口座種類" variant="outlined" onChange={handleChange("bank_account_type")}/>
            <p style={{paddingTop: "50px", paddingBottom: "100px"}}>
              <Button variant="contained" color="primary" type="submit">送信</Button>
            </p>
            </div>
        </div>
      </form>
    </div>
  );
};

export default AccountForm;
