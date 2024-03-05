import { useState } from 'react';
import axios from 'axios';
import { useRouter } from 'next/router';

export default function Login() {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const router = useRouter();

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => { // 型注釈を追加
    e.preventDefault();
    try {
    const { data } = await axios.post('http://localhost:8080/users/me', {
        email,
        password
    });
      localStorage.setItem('token', data.token); // 認証トークンをローカルストレージに保存
      router.push('/main'); // メインページにリダイレクト
    } catch (error) {
    console.error(error);
    alert('ログインに失敗しました。');
    }
};

return (
    <form onSubmit={handleSubmit}>
    <input
        typeof='name'
        value={name}
        onChange={(e) => setName(e.target.value)}
        placeholder="Name"
        required
        />
    <input
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        placeholder="Email"
        required
    />
    <input
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        placeholder="Password"
        required
    />
    <button type="submit">ログイン</button>
    </form>
);
}
