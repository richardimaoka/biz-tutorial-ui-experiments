import { GqlHandsonComponent } from "@/app/components/handson/GqlHandsonComponent";
import { graphql } from "@/libs/gql";
import { request } from "graphql-request";

const queryDefinition = graphql(`
  query appTestTutorialColumnsPage($file: String) {
    _test {
      appTestTutorialColumnsPage {
        ...GqlHandsonComponent
      }
    }
  }
`);

interface PageParams {
  searchParams: {};
}

export default async function Page({ searchParams }: PageParams) {
  // const variables = { step: stepNum };
  const data = await request("http://localhost:8080/query", queryDefinition, {
    /* variables*/
  });

  const fragment = data._test?.appTestTutorialColumnsPage;

  return (
    <div style={{ height: "100svh" }}>
      {fragment && <GqlHandsonComponent fragment={fragment} />}
    </div>
  );
}
