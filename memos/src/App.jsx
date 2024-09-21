import MarkdownRenderer from "./component/MarkdownRenderer";

function App() {
  const tweets = [
    {
      id: 1,
      username: "user1",
      content: "This is tweet 1.This is tweet 1.This is tweet 1.This is tweet 1.This is tweet 1.This is tweet 1.This is tweet 1.This is tweet 1.This is tweet 1.This is tweet 1.This is tweet 1.This is tweet 1.This is tweet 1.",
    },
    {
      id: 2,
      username: "user2",
      content: "This is tweet 2.This is tweet 2.This is tweet 2.This is tweet 2.This is tweet 2.This is tweet 2.This is tweet 2.This is tweet 2.This is tweet 2.This is tweet 2.This is tweet 2.This is tweet 2.This is tweet 2.",
    },
    {
      id: 3,
      username: "user3",
      content: `# Try CommonMark

You can try CommonMark here.  This dingus is powered by
[commonmark.js](https://github.com/commonmark/commonmark.js), the
JavaScript reference implementation.

* Follows [CommonMark](https://commonmark.org)
* Optionally follows [GitHub Flavored Markdown](https://github.github.com/gfm/)
* Renders actual React elements instead of using \`dangerouslySetInnerHTML\`
* Lets you define your own components (to render \`MyHeading\` instead of \`'h1'\`)
* Has a lot of plugins

1. item one
2. item two
   - sublist
   - sublist
   \`\`\`js 
   console.log("112312");
   \`\`\`



A paragraph with *emphasis* and **strong importance**.

> A block quote with ~strikethrough~ and a URL: https://reactjs.org.

* Lists
* [ ] todo
* [x] done

A table:

| a | b |
| - | - |
| 123 | 456 |

# React Markdown Example

- Some text
- Some other text

## Subtitle

### Additional info

This is a [link](https://github.com/remarkjs/react-markdown)
`,
    },
  ];

  return (
    <>
    {tweets.map(tweet => (
      <div key={tweet.id} className="max-w-screen-sm mx-auto bg-white border border-gray-10 p-4 shadow-sm">
        <div className="flex items-start">
          <div className="ml-4">
            <div className="flex items-center">
              <span className="font-bold">{tweet.username}</span>
              <span className="text-gray-500 ml-2">@{tweet.id} · 28分钟</span>
            </div>
            <div className="">
              <MarkdownRenderer content={tweet.content} /> 
            </div>
          </div>
        </div>
      </div>
      ))}
    </>
  );
}

export default App;
