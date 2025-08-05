import Cookies from "js-cookie";
import { useEffect, useState } from "react";
import apiClient from "../utils/apiClient";
import { Button } from "./ui/button";
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
          // Optionally handle token error
        });
    }
  }, []);

  const handleLogout = () => {
    Cookies.remove("accessToken");
    setUserEmail("");
    window.location.href = "/login";
  };

  return (
    <div className="flex flex-col min-h-screen bg-gray-100 dark:bg-gray-900 text-gray-900 dark:text-gray-100 transition-colors">
      <nav className="flex items-center justify-between px-8 py-4 border-b border-gray-300 dark:border-gray-700">
        <div className="flex items-center">
          <span className="font-bold text-xl">Unison</span>
        </div>
        <div className="flex items-center gap-4">
          <span className="text-sm">{userEmail}</span>
          {userEmail ? (
            <Button variant="outline" onClick={handleLogout}>
              Logout
            </Button>
          ) : (
            <Button
              variant="outline"
              onClick={() => {
                window.location.href = "/login";
              }}
            >
              Sign In
            </Button>
          )}
        </div>
      </nav>

      <main className="w-full flex items-center justify-center  pt-4 h-full">
        <div className="flex w-full h-full items-start justify-center">
          <div className="flex  h-full overflow-hidden flex-col p-4 w-full max-w-6xl bg-grey-50 dark:bg-gray-800 rounded-lg shadow-md">
            {/* <div className="flex items-center justify-start">
              <div className="border-1 w-full p-2 rounded-md">toolbar</div>
            </div> */}
            <div className="flex-1 h-full overflow-hidden flex">
              {/* <Editor /> */}
              <form className="flex-1 w-full p-2">
                <textarea
                  className="w-full h-full p-2 border-1 rounded-md resize-none"
                  placeholder="Type something..."
                ></textarea>
              </form>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
};

export default Home;
