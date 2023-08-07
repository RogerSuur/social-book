import { Routes, Route, BrowserRouter } from "react-router-dom";
import NavigationBar from "./components/NavigationBar";
import Profile from "./pages/ProfilePage";
import PostsPage from "./pages/PostsPage";
import Post from "./components/Post";
import Signup from "./pages/SignupPage";
import Login from "./pages/LoginPage";
import NoPage from "./pages/NoPage";
import RequireAuth from "./components/RequireAuth";
import ProfileInfo from "./components/ProfileInfo";
import GroupPage from "./pages/GroupPage";
//import RequireGuest from "./components/RequireGuest";
import Logout from "./components/Logout";
import EventPage from "./pages/EventPage.js";

import "./style.css";

function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<NavigationBar />}>
            <Route path="login" element={<Login />} />
            <Route path="signup" element={<Signup />} />

            <Route element={<RequireAuth />}>
              <Route path="profile" element={<Profile />} />
              <Route path="profile/:id" element={<ProfileInfo />} />
              <Route path="groups/:id" element={<GroupPage />} />
              <Route path="event/:id" element={<EventPage />} />
              <Route path="posts" element={<PostsPage url={"/feedposts"} />} />
              <Route path="posts/:id" element={<Post />} />
              <Route path="logout" element={<Logout />} />
              <Route path="*" element={<NoPage />} />
            </Route>
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
