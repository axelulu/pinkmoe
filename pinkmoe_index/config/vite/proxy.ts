import {
  API_BASE_URL,
  API_TARGET_URL,
  MOCK_API_BASE_URL,
  MOCK_API_TARGET_URL,
  API_UPLOAD_TARGET_URL,
  API_UPLOAD_URL,
  API_SOCKET_URL,
  API_SOCKET_TARGET_URL,
} from '../../config/constant';
import { ProxyOptions } from 'vite';
type ProxyTargetList = Record<string, ProxyOptions>;

const init: ProxyTargetList = {
  // test
  [API_BASE_URL]: {
    target: API_TARGET_URL,
    changeOrigin: true,
    rewrite: (path) => path.replace(new RegExp(`^${API_BASE_URL}`), ''),
  },
  [API_SOCKET_URL]: {
    target: API_SOCKET_TARGET_URL,
    ws: true,
    changeOrigin: true,
    rewrite: (path) => path.replace(new RegExp(`^${API_SOCKET_URL}`), ''),
  },
  [API_UPLOAD_URL]: {
    target: API_UPLOAD_TARGET_URL,
    changeOrigin: true,
    rewrite: (path) => path.replace(new RegExp(`^${API_UPLOAD_URL}`), ''),
  },
  // mock
  [MOCK_API_BASE_URL]: {
    target: MOCK_API_TARGET_URL,
    changeOrigin: true,
    rewrite: (path) => path.replace(new RegExp(`^${MOCK_API_BASE_URL}`), '/api'),
  },
};

export default init;
