import { GqlColumnWrapper } from "@/app/components/tutorial/column/GqlColumnWrapper";
import { graphql } from "@/libs/gql";
import { request } from "graphql-request";

const queryDefinition = graphql(`
  query appTestTutorialColumnsPage {
    _test {
      appTestTutorialColumnsPage {
        columns {
          ...GqlColumnWrapper
        }
      }
    }
  }
`);

interface PageParams {
  searchParams: {};
}

export default async function Page({ searchParams }: PageParams) {
  // const variables = { step: stepNum };
  const data = await request(
    "http://localhost:8080/query",
    queryDefinition
    // variables
  );

  const columns = data._test?.appTestTutorialColumnsPage?.columns;

  if (!columns || columns.length < 1) {
    return <>no data</>;
  }

  const col = columns[0];

  return (
    <div style={{ height: "95svh" }}>
      {col && <GqlColumnWrapper fragment={col} />}
    </div>
  );
}
