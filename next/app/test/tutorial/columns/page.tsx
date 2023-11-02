import { GqlColumnWrapper } from "@/app/components/tutorial/column/GqlColumnWrapper";
import { graphql } from "@/libs/gql";
import { request } from "graphql-request";

const queryDefinition = graphql(`
  query appTestTutorialColumnsPage {
    _test {
      appTestTutorialColumnsPage {
        ...GqlColumnWrapper
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

  const fragment = data._test?.appTestTutorialColumnsPage;

  if (!fragment) {
    return <>no data</>;
  }

  return (
    <div style={{ height: "95svh" }}>
      <GqlColumnWrapper fragment={fragment} />
    </div>
  );
}
