import React, { useState, useEffect } from "react";
import { useLocation } from "react-router-dom";
import {Paper, List, ListItem, Divider} from "@material-ui/core"

import StudentLesson from "./../components/StudentLesson";
import { getOwnerWebsite, getOwnerLessons } from "./../api";

interface WebsiteInfo {
  id: number;
  owner_id: number;
  title: string;
  profile: string;
  theme: string;
  content: string;
}

interface Lesson {
  id: number;
  owner_id: number;
  start_time: number;
  end_time: number;
  meeting_id: string;
  price: number;
}

const ParticipantWeb: React.FC = () => {
  const [websiteInfo, setWebsiteInfo] = useState<WebsiteInfo[]>([]);
  const [lessons, setLessons] = useState<Lesson[]>([]);
  const [isLoading, setIsLoading] = useState(true);

  const path = useLocation().pathname;
  const owner_id: string = path.split("/")[2];

  const containerStyle = {
    textAlign: 'center' as const,
    paddingTop: "100px",
  }

  const innerStyle = {
    paddingTop: "100px",
  }

  const lessonListStyle = {
    paddingTop: "50px",
  }

  useEffect(() => {
    const f = async () => {
      const ownerWebsiteInfo = await getOwnerWebsite(owner_id);
      const ownerLessons = await getOwnerLessons(owner_id);
      if (ownerWebsiteInfo === "NotFound") {
        window.location.href = "/notfound";
      }
      setIsLoading(false);
      setWebsiteInfo(ownerWebsiteInfo);
      setLessons(ownerLessons);
    };
    f();
  }, []);

  if (isLoading) {
    return <p>読み込み中</p>;
  } else if (!websiteInfo[0]) {
    return null;
  } else {
    return (
      <div style={containerStyle}>
        <h1>{websiteInfo[0].title}</h1>
        <div style={innerStyle}>
          <h3>プロフィール</h3>
          <p>{websiteInfo[0].profile}</p>
          <h3>概要</h3>
          <p>{websiteInfo[0].content}</p>
          <div style={lessonListStyle}>
            <h3>受付中のレッスン</h3>
            <Paper>
              <List>
                {lessons.map((lesson, idx) => {
                  return (
                    <div key={lesson.id}>
                      { idx > 0 ? <Divider /> : null }
                      <ListItem>
                        <StudentLesson key={lesson.id} {...lesson} />
                      </ListItem>
                    </div>
                  )
                })}
              </List>
            </Paper>
          </div>
        </div>
      </div>
    );
  }
};

export default ParticipantWeb;
