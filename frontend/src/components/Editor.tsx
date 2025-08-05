import { useState } from "react";
import { createEditor } from "slate";

import { Slate, Editable, withReact } from "slate-react";

const Editor = () => {
  const [editor] = useState(() => withReact(createEditor()));
  const initialValue = [{ type: "paragraph", children: [{ text: "" }] }];
  return (
    <div className="flex-1 w-full p-2 ">
      <Slate editor={editor} initialValue={initialValue}>
        <Editable
          placeholder="Type something..."
          className="p-2 rounded-md w-full h-full"
          style={{
            width: "100%",
            height: "100%",
            border: "none",
            outline: "none",
            boxShadow: "none",
          }}
        />
      </Slate>
    </div>
  );
};

export default Editor;
