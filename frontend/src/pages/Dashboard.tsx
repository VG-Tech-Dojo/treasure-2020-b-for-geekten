import React, { useState, useEffect } from "react";
import {Paper, List, ListItem, Divider} from "@material-ui/core"

import Lesson from "./../components/Lesson";
import useUser from "../hooks/useUser";
import { getAllLesson, getOwnerMe } from "../api";

interface Lesson {
  id: number;
  owner_id: number;
  start_time: number;
  end_time: number;
  meeting_id: string;
  price: number;
}

interface Owner {
  id: number;
  first_name: string;
  last_name: string;
  postal_number: string;
  prefecture: string;
  city: string;
  address: string;
  address_optional: string;
  phone_number: string;
  email: string;
  firebase_uid: string;
  created_at: string;
  updated_at: string;
}

const Dashboard = () => {
  const { user } = useUser();
  const [userToken, setUserToken] = useState("");
  const [lessons, setLessons] = useState<Lesson[]>([]);
  const [owner, setOwner] = useState<Owner>();
  const [webURL, setWebURL] = useState("");

  console.log("user")
  console.log(user)

  useEffect(() => {
    const f = async () => {
      if (!!user) {
        const token = await user!.getIdToken();
        setUserToken(token);

        try {
          const lessons: Lesson[] = await getAllLesson(token);
          console.log(lessons)
          setLessons(lessons)

          const owner: Owner = await getOwnerMe(token)
          console.log(owner)
          setOwner(owner)

          setWebURL(location.href.split("/").slice(0, 3).join("/") + "/owner/" + owner.id + "/web");
        } catch(e) {
          console.error(e)
        }
      }
    }
    f();
  }, [user])

  const paddingStyle = {
    paddingLeft: "50px",
  }

  return (
    <div>
      <h1 style={paddingStyle}>ダッシュボード</h1>
      <h2 style={paddingStyle}>あなたのWebサイトのURL</h2>
      <a style={paddingStyle}href={webURL} target="_blank">{webURL}</a>
      <h2 style={paddingStyle}>登録したレッスン一覧</h2>
      <Paper>
        <List>
          {lessons.map((lesson, idx) => {
            return (
              <div key={lesson.id}>
                { idx > 0 ? <Divider /> : null }
                <ListItem>
                  <Lesson {...lesson} />
                </ListItem>
              </div>
            )
          })}
        </List>
      </Paper>
    </div>
  );
};

export default Dashboard;
