import { useState } from "react";
import { createEditor } from "slate";

import { Slate, Editable, withReact } from "slate-react";

const Editor = () => {
  const [editor] = useState(() => withReact(createEditor()));
  const initialValue = [{ type: "paragraph", children: [{ text: "" }] }];
  return (
    <div>
      <Slate editor={editor} initialValue={initialValue}>
        <Editable
          placeholder="Type something..."
          className="p-4 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
      </Slate>
      <div className="mt-4">
        <button className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">
          Save
        </button>
      </div>
    </div>
  );
};

export default Editor;
