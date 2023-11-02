import { GqlColumnWrappers } from "@/app/components/tutorial/column/GqlColumnWrappers";
import { graphql } from "@/libs/gql";
import { request } from "graphql-request";

const queryDefinition = graphql(`
  query appTestTutorialColumnsPage {
    _test {
      appTestTutorialColumnsPage {
        ...GqlColumnWrappers
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

  return (
    <div style={{ height: "95svh" }}>
      {fragment && <GqlColumnWrappers fragment={fragment} />}
    </div>
  );
}
