import { useQuery } from "@apollo/client";
import { css } from "@emotion/react";
import { Header } from "../components/Header";
import { graphql } from "../libs/gql";

const Home2_Query = graphql(/* GraphQL */ `
  query Home2_Query {
    terminal {
      name
      currentDirectory
      elements {
        __typename
        ... on TerminalCommand {
          ...TerminalCommand_Fragment
        }
        ... on TerminalOutput {
          ...TerminalOutput_Fragment
        }
      }
    }
  }
`);

export default function Home2() {
  const { loading, error, data } = useQuery(Home2_Query);
  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error : {error.message}</p>;

  return (
    data && (
      <>
        <Header />
        <main
          css={css`
            background-color: #333333;
          `}
        >
          <div
            css={css`
              width: 680px;
              margin: 0 auto;
              background-color: white;
            `}
          >
            <pre
              css={css`
                margin: 0px 0px;
                padding: 4px;
                background-color: #1e1e1e;
                color: #f1f1f1;
                border-bottom: 1px solid #333333;
              `}
            >
              <code>
                {
                  "protoc \\\n  --go_out=outdir --go_opt=paths=source_relative \\\n  helloworld.proto"
                }
              </code>
            </pre>
            <pre
              css={css`
                margin: 0px 0px;
                padding: 4px;
                background-color: #1e1e1e;
                color: #f1f1f1;
                border-bottom: 1px solid #333333;
              `}
            >
              <code>outdir/: No such file or directory</code>
            </pre>
          </div>
        </main>
      </>
    )
  );
}
