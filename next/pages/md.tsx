import { css } from "@emotion/react";

export default function Page() {
  const sourceCode = `interface RefetchQueriesOptions<
  TCache extends ApolloCache<any>,
  TResult = Promise<ApolloQueryResult<any>>,
> {
  updateCache?: (cache: TCache) => void;
  include?: Array<string | DocumentNode> | "all" | "active";
  onQueryUpdated?: (
    observableQuery: ObservableQuery<any>,
    diff: Cache.DiffResult<any>,
    lastDiff: Cache.DiffResult<any> | undefined,
  ) => boolean | TResult;
  optimistic?: boolean;
}`;

  return (
    <div
      css={css`
        width: 680px;
        /* margin: 0 auto; */
      `}
    >
      <div
        css={css`
          background-color: #252526;

          h1 {
            font-size: 32px;
            font-weight: 700;
            margin: 21px 0px;
          }

          h2 {
            font-size: 24px;
            font-weight: 700;
            margin: 19px 0px;
          }

          h3 {
            font-size: 18px;
            font-weight: 700;
            margin: 18px 0px;
          }

          p {
            font-size: 14px;
            margin: 16px 0px;
          }

          li {
            font-size: 14px;
          }

          code {
            font-family: "Roboto Mono";
            font-weight: 500;
            font-size: 14px;
            line-height: 18px;
            padding: 1px 4px;
            background-color: #3c3c3c;
          }

          pre {
            padding: 8px;
            background-color: #3c3c3c;
          }

          pre > code {
            font-size: 14px;
            line-height: 20px;
          }
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
          Applications that use Apollo Client require two top-level
          dependencies:
        </p>
        <ul>
          <li>
            <code>@apollo/client</code>: This single package contains virtually
            everything you need to set up Apollo Client. It includes the
            in-memory cache, local state management, error handling, and a
            React-based view layer.
          </li>
          <li>
            <code>graphql</code>: This package provides logic for parsing
            GraphQL queries.
          </li>
        </ul>
        <p>Run the following command to install both of these packages:</p>
        <h3>npm</h3>
        <pre>
          <code>{sourceCode}</code>
        </pre>
        <img src="https://camo.qiitausercontent.com/c6718430970d7f1756e60ef460da266a5476599a/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e616d617a6f6e6177732e636f6d2f302f34353631372f30313562643035382d376561302d653661352d623963622d3336613466623338653539632e706e67" />
      </div>
    </div>
  );
}
