export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  if (process.env.NODE_ENV !== "production") throw new Error("test");

  return (
    <div>
      <div>this is test layout</div>
      {children}
    </div>
  );
}
