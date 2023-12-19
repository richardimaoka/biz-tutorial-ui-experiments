import { TutorialTitle } from "@/app/components/tutorialtitle/TutotiralTitle";

export default async function Page() {
  const images = [
    {
      src: "/images/nextjs-deploy-cloudrun/nextjs.svg",
      width: 120,
      height: 120,
      caption: "Next.jsを",
    },
    {
      src: "/images/nextjs-deploy-cloudrun/cloud_run.svg",
      width: 120,
      height: 120,
      caption: "Cloud Runにデプロイし",
    },
    {
      src: "/images/nextjs-deploy-cloudrun/github-mark-white.svg",
      width: 120,
      height: 120,
      caption: "CI/CDにGitHub Actionsを利用する",
    },
  ];

  return (
    <div style={{ height: "100svh" }}>
      <TutorialTitle
        title="Next.jsをCloud Runにデプロイし、CI/CDにGitHub Actionsを利用するチュートリアル"
        images={images}
      />
    </div>
  );
}
