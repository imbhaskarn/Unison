import Cookies from "js-cookie";
import { useEffect, useState } from "react";
import apiClient from "../utils/apiClient";
import Editor from "./Editor";

const Home = () => {
  const [userEmail, setUserEmail] = useState<string>("");

  useEffect(() => {
    const accessToken = Cookies.get("accessToken");
    if (accessToken) {
      apiClient
        .get("/auth/user", {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        })
        .then((response) => {
          const data = (response.data as { data: { email: string } }).data;
          setUserEmail(data.email);
        })
        .catch((err) => {
          console.log(err);
          // window.location.href = "/login";
        });
    }
  }, []);

  return (
    <div>
      <nav className="flex items-center justify-between px-8 py-4 bg-gray-100 border-b border-gray-300">
        <div className="flex items-center">
          {/* <img src="/logo192.png" alt="Logo" className="h-10 mr-4" /> */}
          <span className="font-bold text-xl">Unison</span>
        </div>
        <div className="flex items-center gap-4">
          <span>{userEmail}</span>
          {!!userEmail.length ? (
            <button
              className="px-4 py-2 bg-blue-700 text-white rounded border-none cursor-pointer hover:bg-blue-800 transition-colors duration-300"
              onClick={() => {
                // Add logout logic here
                alert("Logged out!");
              }}
            >
              Logout
            </button>
          ) : (
            <button
              className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 transition-colors duration-300"
              onClick={() => {
                window.location.href = "/login";
              }}
            >
              Sign In
            </button>
          )}
        </div>
      </nav>
      <div className="flex justify-center items-start min-h-[80vh] bg-blue-50">
        <div className="flex-1" />
        <div className="flex-2 max-w-3xl w-full p-8">
          <Editor />
        </div>
        <div className="flex-1" />
      </div>
    </div>
  );
};

export default Home;
