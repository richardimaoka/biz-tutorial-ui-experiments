import { GqlFileContentPane } from "@/app/components/sourcecode2/openfile/filecontent/GqlFileContentPane";
import { graphql } from "@/libs/gql";
import { request } from "graphql-request";
import Link from "next/link";

const queryDefinition = graphql(`
  query appTestSourcecodeFilecontentPage {
    _test {
      appTestSourcecodeFilecontentPage {
        ...GqlFileContentPane
      }
    }
  }
`);

interface PageParams {
  searchParams: {
    step?: string;
  };
}

export default async function Page({ searchParams }: PageParams) {
  const stepNum = searchParams.step ? Number(searchParams.step) : 1;
  const nextStep = `${stepNum + 1}`;

  const variables = { step: stepNum };
  const data = await request("http://localhost:8080/query", queryDefinition);

  const fragment = data._test?.appTestSourcecodeFilecontentPage;

  if (!fragment) {
    return <>no data</>;
  }

  return (
    <div style={{ height: "95svh" }}>
      <Link href={`/test/terminal?step=${nextStep}`}>next</Link>
      <GqlFileContentPane fragment={fragment} />
    </div>
  );
}
