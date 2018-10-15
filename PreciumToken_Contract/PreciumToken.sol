pragma solidity ^0.4.24;

import "./PreciumTokenBase.sol";
import "./Ownable.sol";

contract PreciumToken is PreciumTokenBase, Ownable {

    string public name;
    string public symbol;
    uint256 public decimals = 18;
    struct lockInfo {
        uint256 lockQuantity;
        uint lockPeriod;
    }
    mapping (address => lockInfo[]) public tokenLockInfo;
    mapping (address => uint256) public unlockQuantity;
    mapping (address => bool) public lockStatus;
    mapping (address => bool) private FreezedWallet;

    function PreciumToken(uint256 initialSupply, string tokenName, uint256 decimalsToken, string tokenSymbol) public {
        decimals = decimalsToken;
        totalSupply_ = initialSupply * 10 ** uint256(decimals);
        emit Transfer(0, msg.sender, totalSupply_);
        balances_[msg.sender] = totalSupply_;
        name = tokenName;
        symbol = tokenSymbol;
        unlockQuantity[msg.sender] = balances_[msg.sender];
    }

    function transfer(address _to, uint256 _value) public returns (bool) {

        bool transferResult;
        uint256 lockQuantity;
        uint256 lockTotalQuantity;
        uint lockPeriod;

        require(FreezedWallet[msg.sender] == false);
        require(FreezedWallet[_to] == false);
        
        if(lockStatus[msg.sender] == false) {
            transferResult = _transfer(msg.sender, _to, _value);
            if (transferResult == true) {
                unlockQuantity[msg.sender] = unlockQuantity[msg.sender].sub(_value);
                unlockQuantity[_to] = unlockQuantity[_to].add(_value);
            }
        }
        else{
            for(uint i = 0; i < tokenLockInfo[msg.sender].length; i++) {
                lockQuantity = tokenLockInfo[msg.sender][i].lockQuantity;
                lockPeriod = tokenLockInfo[msg.sender][i].lockPeriod;

                if(lockPeriod <= now && lockQuantity != 0) {
                    unlockQuantity[msg.sender] = unlockQuantity[msg.sender].add(lockQuantity);
                    tokenLockInfo[msg.sender][i].lockQuantity = 0;
                    lockQuantity = tokenLockInfo[msg.sender][i].lockQuantity;
                }
                lockTotalQuantity = lockTotalQuantity.add(lockQuantity);
            }
            if(lockTotalQuantity == 0)
                lockStatus[msg.sender] = false;
                    
            require(_value <= unlockQuantity[msg.sender]);
            
            transferResult = _transfer(msg.sender, _to, _value);
            if (transferResult == true) {
                unlockQuantity[msg.sender] = unlockQuantity[msg.sender].sub(_value);
                unlockQuantity[_to] = unlockQuantity[_to].add(_value);
            }
        }
        
        return transferResult;
    }
    
    function transferFrom(address _from, address _to, uint256 _value) public returns (bool) {
        
        bool transferResult;
        uint256 lockQuantity;
        uint256 lockTotalQuantity;
        uint lockPeriod;
        
        require(FreezedWallet[_from] == false);
        require(FreezedWallet[_to] == false);
        
        if(lockStatus[_from] == false) {
            transferResult = _transferFrom(_from, _to, _value);
            if (transferResult == true) {
                unlockQuantity[_from] = unlockQuantity[_from].sub(_value);
                unlockQuantity[_to] = unlockQuantity[_to].add(_value);
            }
        }
        else{
            for(uint i = 0; i < tokenLockInfo[_from].length; i++) {
                lockQuantity = tokenLockInfo[_from][i].lockQuantity;
                lockPeriod = tokenLockInfo[_from][i].lockPeriod;

                if(lockPeriod <= now && lockQuantity != 0) {
                    unlockQuantity[_from] = unlockQuantity[_from].add(lockQuantity);
                    tokenLockInfo[_from][i].lockQuantity = 0;
                    lockQuantity = tokenLockInfo[_from][i].lockQuantity;
                }
                lockTotalQuantity = lockTotalQuantity.add(lockQuantity);
            }
            if(lockTotalQuantity == 0)
                lockStatus[_from] = false;
                    
            require(_value <= unlockQuantity[_from]);
            
            transferResult = _transferFrom(_from, _to, _value);
            if (transferResult == true) {
                unlockQuantity[_from] = unlockQuantity[_from].sub(_value);
                unlockQuantity[_to] = unlockQuantity[_to].add(_value);
            }
        }
        
        return transferResult;
    }

    function transferAndLock(address _to, uint256 _value, uint _lockPeriod) onlyOwner public {
       
        bool transferResult;
        
        require(FreezedWallet[_to] == false);
        
        transferResult = _transfer(msg.sender, _to, _value);
        if (transferResult == true) {
            lockStatus[_to] = true;
            tokenLockInfo[_to].push(lockInfo(_value, now + _lockPeriod * 1 days ));
            unlockQuantity[msg.sender] = unlockQuantity[msg.sender].sub(_value);
        }
    }

    function changeLockPeriod(address _owner, uint256 _index, uint _newLockPeriod) onlyOwner public {
        
        require(_index < tokenLockInfo[_owner].length);
        
        tokenLockInfo[_owner][_index].lockPeriod = now + _newLockPeriod * 1 days;
    }
    
    function freezingWallet(address _owner) onlyOwner public {
        
        FreezedWallet[_owner] = true;
    }
    
    function unfreezingWallet(address _owner) onlyOwner public {
        
        FreezedWallet[_owner] = false;
    }
    
    function burn(uint256 _amount) onlyOwner public {
        
        _burn(msg.sender, _amount);
        unlockQuantity[msg.sender] = unlockQuantity[msg.sender].sub(_amount);
    }

    function getNowTime() view public returns(uint res) {
        
        return now;
    }

    function getLockInfo(address _owner, uint256 _index) view public returns(uint256, uint) {
        
        return (tokenLockInfo[_owner][_index].lockQuantity, tokenLockInfo[_owner][_index].lockPeriod);
    }

    function getUnlockQuantity(address _owner) view public returns(uint res) {
        
        return unlockQuantity[_owner];
    }
    
    function getLockStatus(address _owner) view public returns(bool res) {
        
        return lockStatus[_owner];
    }
    
    function getLockCount(address _owner) view public returns(uint res) {
        
        return tokenLockInfo[_owner].length;
    }
    
    function getFreezingInfo(address _owner) view public returns(bool res) {
        
        return FreezedWallet[_owner];
    }
}