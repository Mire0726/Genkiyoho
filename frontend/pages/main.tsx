import { useEffect } from 'react';
import { useRouter } from 'next/router';

export default function Main() {
  const router = useRouter();

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (!token) {
      router.push('/login'); // トークンがなければログインページにリダイレクト
    }
    // ここでトークンを使ってユーザー認証を確認するAPIリクエストを送ることも可能
  }, []);

  return (
    <div>
      <h1>メインページ</h1>
      {/* メインコンテンツ */}
    </div>
  );
}
