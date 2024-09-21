import React from "react";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { Prism as SyntaxHighlighter } from "react-syntax-highlighter";
import { dracula } from "react-syntax-highlighter/dist/esm/styles/prism";
import rehypeRaw from "rehype-raw";

const MarkdownRenderer = ({ content }) => {
  return (
    <ReactMarkdown
      children={content}
      remarkPlugins={[[remarkGfm]]}
      rehypePlugins={[rehypeRaw]} // 启用 HTML 渲染
      components={{
        h1: ({ node, ...props }) => <h1 className="text-2xl font-bold mt-4 mb-2" {...props} />,
        h2: ({ node, ...props }) => <h2 className="text-xl font-bold mt-4 mb-2" {...props} />,
        h3: ({ node, ...props }) => <h3 className="text-lg font-bold mt-4 mb-2" {...props} />,
        p: ({ node, ...props }) => <p className="text-base leading-6 mb-2" {...props} />,
        a: ({ node, ...props }) => <a className="text-blue-500 hover:underline" {...props} />,
        ul: ({ node, ...props }) => <ul className="list-disc list-inside mb-2" {...props} />,
        ol: ({ node, ...props }) => <ol className="list-decimal list-inside mb-2" {...props} />,
        li: ({ node, ...props }) => <li className="mb-1" {...props} />,
        blockquote: ({ node, ...props }) => <blockquote className="border-l-4 border-gray-300 pl-4 italic mb-2" {...props} />,
        // Use a component instead of hrs
        hr(props) {
          const { node, ...rest } = props;
          return <MyFancyRule {...rest} />;
        },
        code(props) {
          const { children, className, node, ...rest } = props;
          const match = /language-(\w+)/.exec(className || "");
          return match ? (
            <SyntaxHighlighter
              {...rest}
              PreTag="div"
              children={String(children).replace(/\n$/, "")}
              language={match[1]}
              style={dracula}
            />
          ) : (
            <code {...rest} className={className}>
              {children}
            </code>
          );
        },
      }}
    ></ReactMarkdown>
  );
};

export default MarkdownRenderer;
