import { GqlTutorialComponent } from "@/app/components/tutorial/GqlTutorialComponent";
import { graphql } from "@/libs/gql";
import { request } from "graphql-request";

const queryDefinition = graphql(`
  query appTestTutorialColumnsPage {
    _test {
      appTestTutorialColumnsPage {
        ...GqlTutorialComponent
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
    <div style={{ height: "100svh" }}>
      {fragment && (
        <GqlTutorialComponent fragment={fragment} selectTab="Terminal" />
      )}
    </div>
  );
}
