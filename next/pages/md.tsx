import { css } from "@emotion/react";

export default function Page() {
  return (
    <div
      css={css`
        width: 680px;
        /* margin: 0 auto; */
      `}
    >
      <h1>Get started with Apollo Client</h1>
      <p>
        Hello! 👋 This short tutorial gets you up and running with Apollo
        Client.
      </p>
      <blockquote>
        <p>
          For an introduction to the entire Apollo platform,{" "}
          <a href="https://www.apollographql.com/tutorials/?utm_source=apollo_docs&amp;utm_medium=referral">
            check out Odyssey, Apollo’s interactive learning platform
          </a>
          .
        </p>
      </blockquote>
      <h2>Step 1: Setup</h2>
      <p>To start this tutorial, do one of the following:</p>
      <ul>
        <li>
          Create a new React project locally with{" "}
          <a href="https://create-react-app.dev/">Create React App</a>, or
        </li>
        <li>
          Create a new React sandbox on{" "}
          <a href="https://codesandbox.io/">CodeSandbox</a>.
        </li>
      </ul>
      <h2>Step 2: Install dependencies</h2>
      <p>
        Applications that use Apollo Client require two top-level dependencies:
      </p>
      <ul>
        <li>
          <code>@apollo/client</code>: This single package contains virtually
          everything you need to set up Apollo Client. It includes the in-memory
          cache, local state management, error handling, and a React-based view
          layer.
        </li>
        <li>
          <code>graphql</code>: This package provides logic for parsing GraphQL
          queries.
        </li>
      </ul>
      <p>Run the following command to install both of these packages:</p>
      <h3>npm</h3>
    </div>
  );
}
